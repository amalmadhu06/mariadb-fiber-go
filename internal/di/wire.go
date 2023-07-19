//go:build wireinject
// +build wireinject

package di

import (
	"github.com/amalmadhu06/mariadb-fiber-go/internal/repository"
	"github.com/amalmadhu06/mariadb-fiber-go/internal/services"
	"github.com/amalmadhu06/mariadb-fiber-go/internal/web"
	"github.com/amalmadhu06/mariadb-fiber-go/internal/web/handler"
	"github.com/amalmadhu06/mariadb-fiber-go/pkg/config"
	"github.com/amalmadhu06/mariadb-fiber-go/pkg/db"
	"github.com/google/wire"
)

func InitializeAPI(cfg config.Config) (*web.ServerHTTP, error) {
	wire.Build(
		// database connection
		db.ConnectDatabase,

		//handler
		handler.NewUserHandler,

		//services
		services.NewUserUsecase,

		//repository
		repository.NewUserRepository,

		//server
		web.NewServerHTTP,
	)

	return &web.ServerHTTP{}, nil
}
