package restapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/Shopify/sarama"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/jackc/pgx/v5/pgxpool"

	"github/user-manager/config"
	errorpkg "github/user-manager/internal/error"
	"github/user-manager/internal/generated/server/models"
	"github/user-manager/internal/generated/server/restapi/operations"
	"github/user-manager/internal/generated/server/restapi/operations/healthcheck"
	"github/user-manager/internal/generated/server/restapi/operations/user"
	hcheckpkg "github/user-manager/internal/infrastructure/handler/healthcheck"
	usrhdlr "github/user-manager/internal/infrastructure/handler/user"
	"github/user-manager/internal/infrastructure/producer/kafka"
	"github/user-manager/internal/infrastructure/producer/kafka/useraction"
	countryrepo "github/user-manager/internal/infrastructure/respository/country"
	usrrepo "github/user-manager/internal/infrastructure/respository/user"
	usrsrv "github/user-manager/internal/service/user"
	loggerpkg "github/user-manager/tools/logger"
)

func configureAPI(api *operations.UserManagerAPI, cfg config.Service) http.Handler {
	// configure the api here
	api.ServeError = serveError

	logger, err := loggerpkg.CreateLogger(cfg.Environment)
	if err != nil {
		panic(err)
	}

	api.Logger = func(s string, i ...interface{}) {
		logger.Info(fmt.Sprintf(s, i))
	}

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	configureHandlers(api, cfg, logger)

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return panicRecoveryMiddleware(
		loggerMiddleware(
			setupGlobalMiddleware(api.Serve(setupMiddlewares)),
			logger,
		),
		logger,
	)
}

func configureHandlers(api *operations.UserManagerAPI, cfg config.Service, logger loggerpkg.ILogger) {
	// -------------
	// repositories
	// -------------

	masterConn, err := connectToPostgres(cfg.Postgres.Master)
	if err != nil {
		logger.WithError(err).Fatal("fail to connect to postgres master")
	}
	slaveConn, err := connectToPostgres(cfg.Postgres.Master)
	if err != nil {
		logger.WithError(err).Fatal("fail to connect to postgres slave")
	}

	userRepo := usrrepo.NewRepository(masterConn, slaveConn)
	countryRepo := countryrepo.NewRepository(slaveConn)

	// -------------
	// producers
	// -------------

	saramaClient, err := sarama.NewClient([]string{fmt.Sprintf("%s:%d", cfg.Kafka.Host, cfg.Kafka.Port)}, nil)
	if err != nil {
		logger.WithError(err).Fatal("fail to create samara client")
	}
	saramaProducer, err := sarama.NewAsyncProducerFromClient(saramaClient)
	if err != nil {
		logger.WithError(err).Fatal("fail to create samara producer")
	}
	eventProducer := kafka.NewProducer(saramaProducer)
	userEventProducer := useraction.NewProducer(eventProducer)

	// -------------
	// services
	// -------------

	userSrv := usrsrv.NewService(userRepo, countryRepo, userEventProducer)

	// -------------
	// handlers
	// -------------

	userHdlr := usrhdlr.NewHandler(userSrv)
	healthcheckHdlr := hcheckpkg.NewHandler(
		func() (string, bool) {
			return "postgres master", masterConn.Ping(context.Background()) == nil
		},
		func() (string, bool) {
			return "postgres slave", slaveConn.Ping(context.Background()) == nil
		},
		func() (string, bool) {
			return "kafka", len(saramaClient.Brokers()) > 0
		},
	)

	api.HealthcheckGetHealthCheckHandler = healthcheck.GetHealthCheckHandlerFunc(healthcheckHdlr.HealthCheck)
	api.UserCreateUserHandler = user.CreateUserHandlerFunc(userHdlr.CreateUser)
	api.UserDeleteUserHandler = user.DeleteUserHandlerFunc(userHdlr.DeleteUser)
	api.UserUpdateUserHandler = user.UpdateUserHandlerFunc(userHdlr.UpdateUser)
	api.UserGetUsersByFiltersHandler = user.GetUsersByFiltersHandlerFunc(userHdlr.GetUsersByFilters)
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}

func connectToPostgres(dsn string) (*pgxpool.Pool, error) {
	pgxConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("fail to parse postgres master config: %w", err)
	}
	conn, err := pgxpool.New(context.Background(), pgxConfig.ConnString())
	if err != nil {
		return nil, fmt.Errorf("fail to connect to postgres master: %w", err)
	}
	err = conn.Ping(context.Background())
	if err != nil {
		return nil, fmt.Errorf("fail to ping postgres master: %w", err)
	}

	return conn, nil
}

func serveError(rw http.ResponseWriter, r *http.Request, err error) {
	rw.Header().Set("Content-Type", "application/json")
	switch e := err.(type) {
	case *errors.CompositeError:
		er := flattenComposite(e)
		// strips composite errors to first element only
		if len(er.Errors) > 0 {
			serveError(rw, r, er.Errors[0])
		} else {
			// guard against empty CompositeError (invalid construct)
			serveError(rw, r, nil)
		}
	case *errors.MethodNotAllowedError:
		rw.Header().Add("Allow", strings.Join(e.Allowed, ","))
		rw.WriteHeader(http.StatusMethodNotAllowed)
		if r == nil || r.Method != http.MethodHead {
			_, _ = rw.Write(errorAsJSON(e.Error(), errorpkg.GetMethodNotAllowedErrCode()))
		}
	case errors.Error:
		value := reflect.ValueOf(e)
		if value.Kind() == reflect.Ptr && value.IsNil() {
			rw.WriteHeader(http.StatusInternalServerError)
			_, _ = rw.Write(errorAsJSON("Unknown error", errorpkg.GetInternalServiceErrCode()))
			return
		}
		rw.WriteHeader(http.StatusUnprocessableEntity)
		if r == nil || r.Method != http.MethodHead {
			_, _ = rw.Write(errorAsJSON(e.Error(), errorpkg.GetValidationErrCode()))
		}
	case nil:
		rw.WriteHeader(http.StatusInternalServerError)
		_, _ = rw.Write(errorAsJSON("Unknown error", errorpkg.GetInternalServiceErrCode()))
	default:
		rw.WriteHeader(http.StatusInternalServerError)
		if r == nil || r.Method != http.MethodHead {
			_, _ = rw.Write(errorAsJSON(e.Error(), errorpkg.GetInternalServiceErrCode()))
		}
	}
}

func errorAsJSON(message string, code string) []byte {
	//nolint:errchkjson
	b, _ := json.Marshal(models.Error{
		Code:    code,
		Message: message,
	})

	return b
}

func flattenComposite(errs *errors.CompositeError) *errors.CompositeError {
	var res []error
	for _, er := range errs.Errors {
		switch e := er.(type) {
		case *errors.CompositeError:
			if len(e.Errors) > 0 {
				flat := flattenComposite(e)
				if len(flat.Errors) > 0 {
					res = append(res, flat.Errors...)
				}
			}
		default:
			if e != nil {
				res = append(res, e)
			}
		}
	}
	return errors.CompositeValidationError(res...)
}
