package controllers

import (
	"geoproperty_be/domain"

	"github.com/gofiber/fiber/v2"
	"github.com/spatial-go/geoos/space"
)

type AreaController struct {
	AreaUsecase domain.AreaUsecase
}

func NewAreaController(au domain.AreaUsecase) *AreaController {
	return &AreaController{
		AreaUsecase: au,
	}
}

func (ac *AreaController) CheckOverlaps(c *fiber.Ctx) error {
	// Get Request Body
	var area struct {
		Geom space.Polygon `json:"geom"`
	}

	if err := c.BodyParser(&area); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// Validate Request Body
	if area.Geom.IsEmpty() {
		return c.Status(fiber.StatusBadRequest).JSON(domain.Response{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid Geometry",
		})
	}

	// Check Overlaps
	overlaps, err := ac.AreaUsecase.Overlaps(area.Geom)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			Status:  fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	// Return Response
	return c.Status(fiber.StatusOK).JSON(domain.Response{
		Status:  fiber.StatusOK,
		Message: "Success",
		Data: struct {
			Overlaps bool `json:"overlaps"`
		}{
			Overlaps: overlaps,
		},
	})
}
