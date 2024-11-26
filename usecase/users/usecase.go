package users

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"geoproperty_be/config"
	"geoproperty_be/domain"
	"geoproperty_be/utils"
	"net/http"
	"net/url"
	"time"
)

type UseCase struct {
	RepositoryUser domain.UsersRepository
}

// RefreshToken implements domain.UsersUseCase.
func (u *UseCase) RefreshToken(token string) (*domain.Token, error) {
	// Extract Token
	dataClaims, err := utils.ValidateToken(token)
	if err != nil {
		return nil, err
	}

	// Check Expired
	exp := dataClaims["exp"].(float64)
	if time.Now().Unix() > int64(exp) {
		return nil, errors.New("token expired")
	}

	// Get Detail User from email
	user, err := u.RepositoryUser.Find(map[string]any{
		"email": dataClaims["email"].(string),
	})
	if err != nil {
		return nil, err
	}

	if len(*user) == 0 {
		return nil, errors.New("user not found")
	}

	// Generate token
	param := map[string]any{
		"id":    (*user)[0].ID,
		"email": (*user)[0].Email,
		"name":  (*user)[0].Name,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token, err = utils.GenerateToken(param)
	if err != nil {
		return nil, err
	}

	// Refresh Token
	refreshToken, err := utils.GenerateToken(map[string]any{
		"email": (*user)[0].Email,
		"exp":   time.Now().Add(time.Hour * 24 * 7).Unix(),
	})
	if err != nil {
		return nil, errors.New("cannot generate refresh token")
	}

	// Return Token
	return &domain.Token{
		Token:        token,
		RefreshToken: refreshToken,
	}, nil
}

// ExtractTokenGoogle implements domain.UsersUseCase.
func (u *UseCase) ExtractTokenGoogle(code string) (string, error) {
	token, err := config.OauthConfGl.Exchange(context.TODO(), code)

	if err != nil {
		return "", err
	}

	response, errors := http.Get("https://www.googleapis.com/oauth2/v1/userinfo?access_token=" + url.QueryEscape(token.AccessToken))

	if errors != nil {
		return "", errors
	}

	defer response.Body.Close()

	// Decode Response Body
	var data map[string]any
	json.NewDecoder(response.Body).Decode(&data)

	fmt.Println(data)

	// Find user by email
	params := map[string]any{
		"email": data["email"],
	}

	users, err := u.RepositoryUser.Find(params)

	if err != nil {
		return "", err
	}

	if len(*users) == 0 {
		// Register User
		user := domain.Users{
			Name:     data["name"].(string),
			Email:    data["email"].(string),
			Password: "",
		}

		user_new, err := u.RepositoryUser.Insert(&user)

		if err != nil {
			return "", err
		}

		// Generate token
		param := map[string]any{
			"id":    user_new.ID,
			"email": user_new.Email,
			"name":  user_new.Name,
			"exp":   time.Now().Add(time.Hour * 24).Unix(),
		}

		token_jwt, err := utils.GenerateToken(param)

		if err != nil {
			return "", err
		}

		return token_jwt, nil
	}

	// Generate token
	param := map[string]any{
		"id":    (*users)[0].ID,
		"email": (*users)[0].Email,
		"name":  (*users)[0].Name,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token_jwt, err := utils.GenerateToken(param)

	if err != nil {
		return "", err
	}

	return token_jwt, nil
}

// Register implements domain.UsersUseCase.
func (u *UseCase) Register(user *domain.Users) (*domain.Users, error) {
	//Check Email Exist
	users, err := u.RepositoryUser.Find(map[string]any{
		"email": user.Email,
	})

	if err != nil {
		return nil, errors.New("cannot find user for register")
	}

	if len(*users) > 0 {
		return nil, errors.New("email already exist")
	}

	// Hash Password
	hashedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		return nil, errors.New("cannot hash password")
	}

	user.Password = hashedPassword

	// Insert User
	user, err = u.RepositoryUser.Insert(user)

	if err != nil {
		return nil, errors.New("cannot insert user")
	}

	return user, nil
}

// Login implements domain.UsersUseCase.
func (u *UseCase) Login(email string, password string) (*domain.Token, error) {
	// Get Detail User
	users, err := u.RepositoryUser.Find(map[string]any{
		"email": email,
	})

	if err != nil {
		return nil, errors.New("cannot find user")
	}

	if len(*users) == 0 {
		return nil, errors.New("email not found")
	}

	// Check Password
	if !utils.CheckPasswordHash(password, (*users)[0].Password) {
		return nil, errors.New("wrong password")
	}

	// Generate Token
	token, err := utils.GenerateToken(map[string]any{
		"id":    (*users)[0].ID,
		"email": (*users)[0].Email,
		"name":  (*users)[0].Name,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	if err != nil {
		return nil, errors.New("cannot generate token")
	}

	// Generate Refresh Token
	refreshToken, err := utils.GenerateToken(map[string]any{
		"email": (*users)[0].Email,
		"exp":   time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	if err != nil {
		return nil, errors.New("cannot generate refresh token")
	}

	return &domain.Token{
		Token:        token,
		RefreshToken: refreshToken,
	}, nil
}

func NewUseCase(repositoryUser domain.UsersRepository) domain.UsersUseCase {
	return &UseCase{
		RepositoryUser: repositoryUser,
	}
}
