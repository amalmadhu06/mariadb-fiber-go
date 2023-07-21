package handler

import (
	"context"
	"github.com/amalmadhu06/mariadb-fiber-go/internal/services/interfaces"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	_ "go.mongodb.org/mongo-driver/bson"
	"net/http"
	"strings"
)

type OfferHandler struct {
	offerUsecase interfaces.OfferUsecase
}

func NewOfferHandler(usecase interfaces.OfferUsecase) *OfferHandler {
	return &OfferHandler{
		offerUsecase: usecase,
	}
}

func (o *OfferHandler) GetOffer(c *fiber.Ctx) error {
	param := c.Params("country")
	country := strings.ToUpper(param)
	// Todo : Pagination
	if country == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Error{
			Message: "country name must be included in query",
		})
	}

	respFormat := c.Get("Accept")

	offers, err := o.offerUsecase.GetOffer(context.TODO(), country)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	if respFormat == "application/bson" {
		_, offerBytes, err := bson.MarshalValue(offers)
		if err != nil {
			return fiber.NewError(http.StatusInternalServerError, "failed to marshal bson data")
		}
		return c.Status(http.StatusOK).Send(offerBytes)
	}
	// keeping json as the default response format
	return c.Status(http.StatusOK).JSON(offers)
}
