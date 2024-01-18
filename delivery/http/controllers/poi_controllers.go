package controllers

import (
	"geoproperty_be/domain"

	"github.com/gofiber/fiber/v2"
	"github.com/spatial-go/geoos/space"
)

type PoiController struct {
	UsesUseCase domain.POIUsecase
}

func NewPoiController(u domain.POIUsecase) *PoiController {
	return &PoiController{
		UsesUseCase: u,
	}
}

// GeoJSON is a function to get geojson.
func (u *PoiController) GeoJSON(c *fiber.Ctx) error {
	var point struct {
		CenterPoint space.Point `json:"center_point"`
	}

	if err := c.BodyParser(&point); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.Response{
			Status:  fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	geoData, err := u.UsesUseCase.GeoJSON(point.CenterPoint)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.Response{
			Status:  fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(geoData)
}
