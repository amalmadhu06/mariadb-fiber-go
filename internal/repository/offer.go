package repository

import (
	"context"
	"errors"
	"github.com/amalmadhu06/mariadb-fiber-go/internal/entities"
	"github.com/amalmadhu06/mariadb-fiber-go/internal/repository/interfaces"
	"gorm.io/gorm"
)

type offerDatabase struct {
	DB *gorm.DB
}

func NewOfferRepository(db *gorm.DB) interfaces.OfferRepo {
	return &offerDatabase{
		DB: db,
	}
}

func (o *offerDatabase) GetOffer(ctx context.Context, country string) ([]entities.OfferCompany, error) {
	var latestOffer []entities.OfferCompany

	query := `SELECT * FROM offer_company WHERE country = ?;`

	err := o.DB.Raw(query, country).Scan(&latestOffer).Error
	if err != nil {
		return nil, err
	}

	if latestOffer == nil {
		return nil, errors.New("no data found")
	}
	return latestOffer, nil
}
