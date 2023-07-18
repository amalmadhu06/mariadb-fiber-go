package repository

import (
	"github.com/amalmadhu06/mariadb-fiber-go/pkg/repository/interfaces"
	"gorm.io/gorm"
)

type userDatabase struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) interfaces.UserRepo {
	return &userDatabase{DB}
}
