//go:build wireinject
// +build wireinject

package di

import (
	"github.com/amalmadhu06/mariadb-fiber-go/pkg/config"
	"github.com/amalmadhu06/mariadb-fiber-go/pkg/db"
	"github.com/amalmadhu06/mariadb-fiber-go/pkg/http"
	"github.com/amalmadhu06/mariadb-fiber-go/pkg/http/handler"
	"github.com/amalmadhu06/mariadb-fiber-go/pkg/repository"
	"github.com/amalmadhu06/mariadb-fiber-go/pkg/usecase"
	"github.com/google/wire"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(
		// database connection
		db.ConnectDatabase,

		//handler
		handler.NewUserHandler,

		//usecase
		usecase.NewUserUsecase,

		//repository
		repository.NewUserRepository,

		//server
		http.NewServerHTTP,
	)

	return &http.ServerHTTP{}, nil
}
