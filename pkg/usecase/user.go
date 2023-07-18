package usecase

import (
	"github.com/amalmadhu06/mariadb-fiber-go/pkg/repository/interfaces"
	"github.com/amalmadhu06/mariadb-fiber-go/pkg/usecase/services"
)

type userUsecase struct {
	userRepo interfaces.UserRepo
}

func NewUserUsecase(userRepo interfaces.UserRepo) services.UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}
