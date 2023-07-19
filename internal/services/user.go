package services

import (
	"github.com/amalmadhu06/mariadb-fiber-go/internal/repository/interfaces"
	"github.com/amalmadhu06/mariadb-fiber-go/internal/services/interfaces"
)

type userUsecase struct {
	userRepo interfaces.UserRepo
}

func NewUserUsecase(userRepo interfaces.UserRepo) interfaces.UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}
