package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"geoproperty_be/domain"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spatial-go/geoos/space"
)

type PropertyController struct {
	PropertyUseCase domain.PropertyUsecase
	AssetUseCase    domain.AssetUsecase
}

func NewPropertyController(p domain.PropertyUsecase, a domain.AssetUsecase) *PropertyController {
	return &PropertyController{
		PropertyUseCase: p,
		AssetUseCase:    a,
	}
}

// Insert is a function to insert property.
func (p *PropertyController) Insert(ctx *fiber.Ctx) error {
	// Form file
	form, err := ctx.MultipartForm()

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			Status:  fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	context := context.Background()
	files := form.File["files"]
	err = p.AssetUseCase.UploadMultipleAsset(context, files)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			Status:  fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	var property domain.Property[space.Point, space.Polygon]

	data := ctx.FormValue("data")
	err = json.Unmarshal([]byte(data), &property)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			Status:  fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	newProperty, err := p.PropertyUseCase.Insert(property)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			Status:  fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(domain.Response{
		Status:  fiber.StatusOK,
		Message: "success",
		Data:    newProperty,
	})
}

// Get Detail Property
func (p *PropertyController) GetDetail(ctx *fiber.Ctx) error {
	uid := ctx.Params("uid")

	property, err := p.PropertyUseCase.FindDetail(uid)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			Status:  fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(domain.Response{
		Status:  fiber.StatusOK,
		Message: "success",
		Data:    property,
	})
}

// Get All Property
func (p *PropertyController) GetOwn(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	userID := int(claims["id"].(float64))

	properties, err := p.PropertyUseCase.FindAll(map[string]any{
		"user_id": userID,
	})

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			Status:  fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(domain.Response{
		Status:  fiber.StatusOK,
		Message: "success",
		Data:    properties,
	})
}

// Delete Property
func (p *PropertyController) Delete(ctx *fiber.Ctx) error {
	uid := ctx.Params("uid")

	err := p.PropertyUseCase.Delete(uid)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			Status:  fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(domain.Response{
		Status:  fiber.StatusOK,
		Message: "success",
	})
}

// Get By Center Point
func (p *PropertyController) GetByCenterPoint(ctx *fiber.Ctx) error {
	var param struct {
		CenterPoint space.Point `json:"center_point"`
	}

	if err := ctx.BodyParser(&param); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			Status:  fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	fmt.Println(param.CenterPoint)

	properties, err := p.PropertyUseCase.GetByGeom("point", param.CenterPoint, space.Polygon{})

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			Status:  fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(properties)
}

// Get By Polygon
func (p *PropertyController) GetByPolygon(ctx *fiber.Ctx) error {
	var param struct {
		Polygon space.Polygon `json:"polygon"`
	}

	if err := ctx.BodyParser(&param); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			Status:  fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	properties, err := p.PropertyUseCase.GetByGeom("polygon", space.Point{}, param.Polygon)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			Status:  fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(properties)
}

// Get By Prompt
func (p *PropertyController) GetByPrompt(ctx *fiber.Ctx) error {
	var param struct {
		Prompt string `json:"prompt"`
	}

	if err := ctx.BodyParser(&param); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			Status:  fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	properties, err := p.PropertyUseCase.GetPropertyByPrompt(param.Prompt)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			Status:  fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(properties)
}
