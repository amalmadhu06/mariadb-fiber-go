package handler

import (
	"fmt"
	"github.com/amalmadhu06/mariadb-fiber-go/internal/entities"
	"github.com/amalmadhu06/mariadb-fiber-go/internal/services/mockServices"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestOfferHandler_GetOffer(t *testing.T) {
	ctrl := gomock.NewController(t)
	service := mockServices.NewMockOfferUsecase(ctrl)
	handler := NewOfferHandler(service)

	testCases := []struct {
		name             string
		country          string
		buildStub        func(offerSvc mockServices.MockOfferUsecase)
		expectedResponse []entities.OfferCompany
	}{
		{
			name:    "valid country",
			country: "US",
			buildStub: func(offerSvc mockServices.MockOfferUsecase) {
				offerSvc.
					EXPECT().
					GetOffer(gomock.Any(), "US").
					Times(1).
					Return(nil, nil)
			},
			expectedResponse: nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			tt.buildStub(*service)
			app := fiber.New()
			recorder := httptest.NewRecorder()

			app.Get("/offer/:country", handler.GetOffer)

			req := httptest.NewRequest(http.MethodGet, "/offer/US", nil)
			app.Test(req)
			fmt.Println(recorder.Body.Bytes())
			//Todo : complete handler testing
		})
	}
}
