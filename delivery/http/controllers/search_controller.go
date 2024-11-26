package controllers

import (
	"geoproperty_be/domain"

	"github.com/gofiber/fiber/v2"
)

type SearchController struct {
	SearchUseCase domain.SearchUseCase
}

func NewSearchController(u domain.SearchUseCase) *SearchController {
	return &SearchController{
		SearchUseCase: u,
	}
}

// Search is a function to search property.
func (u *SearchController) Search(c *fiber.Ctx) error {
	keyword := c.Params("keyword")

	if keyword == "" {
		return c.Status(fiber.StatusBadRequest).JSON(domain.Response{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid keyword",
		})
	}

	searches, err := u.SearchUseCase.Search(keyword)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.Response{
			Status:  fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	if len(*searches) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(domain.Response{
			Status:  fiber.StatusNotFound,
			Message: "Not Found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(domain.Response{
		Status:  fiber.StatusOK,
		Message: "success",
		Data:    searches,
	})
}

// GetAll is a function to get all property.
func (u *SearchController) GetAll(c *fiber.Ctx) error {
	searches, err := u.SearchUseCase.GetAll()

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.Response{
			Status:  fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	if len(*searches) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(domain.Response{
			Status:  fiber.StatusNotFound,
			Message: "Not Found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(domain.Response{
		Status:  fiber.StatusOK,
		Message: "success",
		Data:    searches,
	})
}
