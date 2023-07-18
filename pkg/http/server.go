package http

import (
	"github.com/amalmadhu06/mariadb-fiber-go/pkg/http/handler"
	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
	"log"
)

type ServerHTTP struct {
	app *fiber.App
}

func NewServerHTTP(userHandler *handler.UserHandler) *ServerHTTP {
	app := fiber.New()

	// Todo : add logger for troubleshooting
	//app.Use()

	//server swagger ui
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	//set up routes based on a base path
	app.Get("/", userHandler.Hello)

	return &ServerHTTP{
		app: app,
	}
}

func (s *ServerHTTP) Start() {
	if err := s.app.Listen(":3000"); err != nil {
		log.Fatal("failed to start http server on port 3000: ", err)
	}
	log.Println("server listening to port 3000")
}
