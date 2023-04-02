package common

import (
	"github.com/go-playground/validator/v10"
)

type Common struct {
	Validate *validator.Validate
}

func New() (*Common, error) {
	validate := getValidator()

	common := Common{}

	common.Validate = validate

	return &common, nil
}
