package controllers

import (
	"geoproperty_be/config"
	"geoproperty_be/domain"
	"geoproperty_be/utils"
	"net/url"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	UsesUseCase domain.UsersUseCase
}

func NewAuthController(u domain.UsersUseCase) *AuthController {
	return &AuthController{
		UsesUseCase: u,
	}
}

// Register is a function to register user.
func (u *AuthController) Register(c *fiber.Ctx) error {
	var user domain.Users

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.Response{
			Status:  fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	// Validate Request Body
	if err := utils.ValidateStruct(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.Response{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid request body",
		})
	}

	users, err := u.UsesUseCase.Register(&user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.Response{
			Status:  fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(domain.Response{
		Status:  fiber.StatusCreated,
		Message: "success",
		Data:    users,
	})
}

// Login is a function to login user.
func (u *AuthController) Login(c *fiber.Ctx) error {
	var auth domain.Auth

	if err := c.BodyParser(&auth); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.Response{
			Status:  fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	token, err := u.UsesUseCase.Login(auth.Email, auth.Password)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.Response{
			Status:  fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(domain.Response{
		Status:  fiber.StatusOK,
		Message: "success",
		Data:    token,
	})
}

// LoginbyGoogle is a function to login user by google.
func (u *AuthController) LoginGoogle(c *fiber.Ctx) error {
	URL, err := url.Parse(config.OauthConfGl.Endpoint.AuthURL)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			Status:  fiber.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}

	// Generate oauth state string
	config.OauthStateStringGl, err = utils.RandomString(32)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			Status:  fiber.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}

	parameters := url.Values{}
	parameters.Add("client_id", config.OauthConfGl.ClientID)
	parameters.Add("scope", strings.Join(config.OauthConfGl.Scopes, " "))
	parameters.Add("redirect_uri", config.OauthConfGl.RedirectURL)
	parameters.Add("response_type", "code")
	parameters.Add("state", config.OauthStateStringGl)
	URL.RawQuery = parameters.Encode()
	url := URL.String()

	return c.Redirect(url, fiber.StatusTemporaryRedirect)
}

// CallbackGoogle is a function to callback user by google.
func (u *AuthController) CallbackGoogle(c *fiber.Ctx) error {
	state := c.FormValue("state")

	if state != config.OauthStateStringGl {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			Status:  fiber.StatusInternalServerError,
			Message: "Invalid state parameter",
			Data:    nil,
		})
	}

	code := c.FormValue("code")

	token, err := u.UsesUseCase.ExtractTokenGoogle(code)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			Status:  fiber.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.Redirect("https://geoproperty.nerdvana-hub.com/auth?token="+token, fiber.StatusTemporaryRedirect)
}

// RefreshToken is a function to refresh token.
func (u *AuthController) RefreshToken(c *fiber.Ctx) error {
	type Token struct {
		Token string `json:"refresh_token" validate:"required"`
	}

	var token Token

	if err := c.BodyParser(&token); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.Response{
			Status:  fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	// Validate Request Body
	if err := utils.ValidateStruct(token); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.Response{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid request body",
		})
	}

	newToken, err := u.UsesUseCase.RefreshToken(token.Token)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.Response{
			Status:  fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(domain.Response{
		Status:  fiber.StatusOK,
		Message: "success",
		Data:    newToken,
	})
}
