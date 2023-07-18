package main

import (
	"github.com/amalmadhu06/mariadb-fiber-go/pkg/config"
	"github.com/amalmadhu06/mariadb-fiber-go/pkg/di"
	"log"
)

func main() {

	// load configuration values
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("failed to load config: ", err)
	}

	// initialize apis using dependencies
	server, err := di.InitializeAPI(cfg)
	if err != nil {
		log.Fatal("failed to start server: ", err)
	}

	// start the server
	server.Start()
}
