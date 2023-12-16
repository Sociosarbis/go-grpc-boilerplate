package common

import (
	"github.com/go-playground/validator/v10"
	"github.com/sociosarbis/grpc/boilerplate/internal/web/req"
	cmdValidator "github.com/sociosarbis/grpc/boilerplate/internal/web/req/validators/cmd"
)

func getValidator() *validator.Validate {
	validate := validator.New()
	validate.RegisterStructValidation(cmdValidator.ValidateCmdItem, req.CmdItem{})
	return validate
}
