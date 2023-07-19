package handler

import (
	"github.com/amalmadhu06/mariadb-fiber-go/internal/services/interfaces"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userUsecase interfaces.UserUsecase
}

func NewUserHandler(usecase interfaces.UserUsecase) *UserHandler {
	return &UserHandler{
		userUsecase: usecase,
	}
}
func (u *UserHandler) Hello(c *fiber.Ctx) error {
	return c.SendString("Hello World!!!")
}
