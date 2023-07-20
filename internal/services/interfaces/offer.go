package interfaces

import (
	"context"
	"github.com/amalmadhu06/mariadb-fiber-go/internal/entities"
)

type OfferUsecase interface {
	GetOffer(ctx context.Context, country string) ([]entities.OfferCompany, error)
}
