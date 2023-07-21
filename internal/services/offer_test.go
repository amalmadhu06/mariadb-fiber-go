package services

import (
	"context"
	"github.com/amalmadhu06/mariadb-fiber-go/internal/entities"
	"github.com/amalmadhu06/mariadb-fiber-go/internal/memory/mockMemory"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOfferUsecase_GetOffer(t *testing.T) {
	ctrl := gomock.NewController(t)
	offerRepo := mockMemory.NewMockOfferMemory(ctrl)
	offerMemory := mockMemory.NewMockOfferMemory(ctrl)

	svc := NewOfferUsecase(offerRepo, offerRepo)

	testData := []struct {
		name           string
		country        string
		buildStub      func(offerMem mockMemory.MockOfferMemory)
		expectedOutput []entities.OfferCompany
	}{
		{
			name:    "valid country",
			country: "US",
			buildStub: func(offerMem mockMemory.MockOfferMemory) {
				offerRepo.EXPECT().
					GetOffer(
						gomock.Any(), "US").Times(1).
					Return([]entities.OfferCompany{
						{OfferID: 1},
					}, nil)
			},
			expectedOutput: []entities.OfferCompany{{OfferID: 1}},
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			tt.buildStub(*offerMemory)
			actualData, err := svc.GetOffer(context.TODO(), tt.country)
			assert.Equal(t, actualData, tt.expectedOutput)
			assert.NoError(t, err)
		})
	}
}
