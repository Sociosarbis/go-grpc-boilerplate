package common

import (
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type Common struct {
	Logger   *zap.Logger
	Validate *validator.Validate
}

func New(logger *zap.Logger) (*Common, error) {
	validate := getValidator()

	common := Common{
		logger,
		validate,
	}

	return &common, nil
}
