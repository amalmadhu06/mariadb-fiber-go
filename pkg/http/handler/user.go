package handler

import (
	"github.com/amalmadhu06/mariadb-fiber-go/pkg/usecase/services"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userUsecase services.UserUsecase
}

func NewUserHandler(usecase services.UserUsecase) *UserHandler {
	return &UserHandler{
		userUsecase: usecase,
	}
}
func (u *UserHandler) Hello(c *fiber.Ctx) error {
	return c.SendString("Hello World!!!")
}
