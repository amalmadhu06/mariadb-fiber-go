package web

import (
	"github.com/amalmadhu06/mariadb-fiber-go/internal/web/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	fiberSwagger "github.com/swaggo/fiber-swagger"
	"log"
)

type ServerHTTP struct {
	app *fiber.App
}

func NewServerHTTP(offerHandler *handler.OfferHandler) *ServerHTTP {

	app := fiber.New()
	// middleware for logging
	app.Use(logger.New())
	// middleware for compressing response
	app.Use(compress.New())

	//server swagger ui
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	//set up routes based on a base path
	app.Get("/offer/:country", offerHandler.GetOffer)

	return &ServerHTTP{
		app: app,
	}
}

func (s *ServerHTTP) Start() {
	log.Println("starting server on port 3000")
	if err := s.app.Listen(":3000"); err != nil {
		log.Fatal("failed to start web server on port 3000: ", err)
	}
}

func (s *ServerHTTP) ShutDown() error {
	log.Println("Shutting down the server gracefully")
	if err := s.app.Shutdown(); err != nil {
		return err
	}
	log.Println("Server shut down gracefully")
	return nil
}
