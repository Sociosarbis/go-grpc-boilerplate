package common

import "github.com/go-playground/validator/v10"

func getValidator() *validator.Validate {
	validate := validator.New()

	return validate
}
