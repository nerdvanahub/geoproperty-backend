package utils

import (
	"os"

	jwt "github.com/dgrijalva/jwt-go"
)

func GenerateToken(data map[string]any) (string, error) {
	var dataClaims jwt.MapClaims = data

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, dataClaims)

	secret := []byte(os.Getenv("JWT_SECRET"))

	tokenString, err := token.SignedString(secret)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (jwt.MapClaims, error) {
	var dataClaims jwt.MapClaims

	secret := []byte(os.Getenv("JWT_SECRET"))

	token, err := jwt.ParseWithClaims(tokenString, &dataClaims, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, err
	}

	return dataClaims, nil
}
