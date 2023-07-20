//go:build wireinject
// +build wireinject

package di

import (
	"github.com/amalmadhu06/mariadb-fiber-go/internal/memory"
	"github.com/amalmadhu06/mariadb-fiber-go/internal/repository"
	"github.com/amalmadhu06/mariadb-fiber-go/internal/services"
	"github.com/amalmadhu06/mariadb-fiber-go/internal/web"
	"github.com/amalmadhu06/mariadb-fiber-go/internal/web/handler"
	"github.com/amalmadhu06/mariadb-fiber-go/pkg/config"
	"github.com/amalmadhu06/mariadb-fiber-go/pkg/db"
	"github.com/google/wire"
)

// InjectDependencies takes in configuration values and inject dependencies
// using google wire
func InjectDependencies(cfg config.Config) (*web.ServerHTTP, error) {
	wire.Build(
		// database connection
		db.ConnectDatabase,

		//handler
		handler.NewOfferHandler,

		//services
		services.NewOfferUsecase,

		// memory
		memory.NewOfferMemory,

		//repository
		repository.NewOfferRepository,

		//server
		web.NewServerHTTP,
	)

	return &web.ServerHTTP{}, nil
}
