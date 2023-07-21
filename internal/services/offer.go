package services

import (
	"context"
	"errors"
	"github.com/amalmadhu06/mariadb-fiber-go/internal/entities"
	mem "github.com/amalmadhu06/mariadb-fiber-go/internal/memory/interfaces"
	repo "github.com/amalmadhu06/mariadb-fiber-go/internal/repository/interfaces"
	svc "github.com/amalmadhu06/mariadb-fiber-go/internal/services/interfaces"
)

type OfferUsecase struct {
	offerRepo   repo.OfferRepo
	offerMemory mem.OfferMemory
}

func NewOfferUsecase(offerRepo repo.OfferRepo, offerMem mem.OfferMemory) svc.OfferUsecase {
	return &OfferUsecase{
		offerRepo:   offerRepo,
		offerMemory: offerMem,
	}
}

func (o *OfferUsecase) GetOffer(ctx context.Context, country string) ([]entities.OfferCompany, error) {

	if country == "US" || country == "FR" || country == "BR" || country == "CA" {
		// if data is present in memory, retrieve and send it
		offer, err := o.offerMemory.GetOffer(ctx, country)
		if err != nil {
			return nil, err
		}

		if offer != nil {
			return offer, nil
		}

		// if data is not present in memory, retrieve it from the database
		return o.offerRepo.GetOffer(ctx, country)
	}
	return nil, errors.New("invalid country name")
}
