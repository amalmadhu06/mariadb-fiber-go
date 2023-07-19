package services

import (
	repo "github.com/amalmadhu06/mariadb-fiber-go/internal/repository/interfaces"
	svc "github.com/amalmadhu06/mariadb-fiber-go/internal/services/interfaces"
)

type userUsecase struct {
	userRepo repo.UserRepo
}

func NewUserUsecase(userRepo repo.UserRepo) svc.UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}
