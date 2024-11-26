package utils

import "github.com/go-playground/validator/v10"

func ValidateStruct(s interface{}) error {
	// Validate Request with validator/v10
	validate := validator.New()

	if err := validate.Struct(s); err != nil {
		return err
	}

	return nil
}
