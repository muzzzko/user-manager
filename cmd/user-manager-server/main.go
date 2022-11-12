package main

import (
	"github/user-manager/config"
	"github/user-manager/internal/generated/server/restapi"
	"github/user-manager/internal/generated/server/restapi/operations"
	"log"

	"github.com/go-openapi/loads"
	"github.com/vrischmann/envconfig"
)

func main() {

	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	// create new service API
	api := operations.NewUserManagerAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	cfg := config.Service{}
	if err := envconfig.InitWithPrefix(&cfg, "user_manager"); err != nil {
		panic(err)
	}

	server.ConfigureAPI(cfg)

	server.Port = cfg.Server.Port
	server.Host = cfg.Server.Host
	server.GracefulTimeout = cfg.Server.GracefulTimeout
	server.ReadTimeout = cfg.Server.ReadTimeout
	server.WriteTimeout = cfg.Server.WriteTimeout

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
