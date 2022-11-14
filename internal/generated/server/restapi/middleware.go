package restapi

import (
	"encoding/json"
	"net/http"
	"runtime/debug"

	errorpkg "github/user-manager/internal/error"
	"github/user-manager/internal/generated/server/models"
	loggerpkg "github/user-manager/tools/logger"
)

func loggerMiddleware(handler http.Handler, logger loggerpkg.ILogger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := loggerpkg.EnrichContext(r.Context(), logger)
		r = r.WithContext(ctx)

		handler.ServeHTTP(w, r)
	})
}

func panicRecoveryMiddleware(handler http.Handler, logger loggerpkg.ILogger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if p := recover(); p != nil {
				logger.WithPanic(string(debug.Stack())).Error("panic recovered")

				w.Header().Add("content-type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)

				err := models.Error{
					Code:    errorpkg.GetInternalServiceErrCode(),
					Message: "something went wrong",
				}
				resp, _ := json.Marshal(err)

				_, _ = w.Write(resp)
			}
		}()

		handler.ServeHTTP(w, r)
	})
}
