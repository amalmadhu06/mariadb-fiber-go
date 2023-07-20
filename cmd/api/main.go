package main

import (
	"github.com/amalmadhu06/mariadb-fiber-go/internal/di"
	"github.com/amalmadhu06/mariadb-fiber-go/pkg/config"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	// loading required configuration values
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("failed to load config: ", err)
	}

	// injecting dependencies and initializing server
	server, err := di.InjectDependencies(cfg)
	if err != nil {
		log.Fatal("failed to start server: ", err)
	}

	// channel for sending server shutdown signal
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)

	// starting server in a separate go routine
	go func() {
		server.Start()
	}()

	// receiving server shutdown signal
	sig := <-signalCh
	log.Println("Received signal : ", sig)

	// graceful shutdown of the server
	if err := server.ShutDown(); err != nil {
		log.Fatalf("Server shutdown failed: %v\n", err)
	}

}
