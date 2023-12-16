package cmd

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/sociosarbis/grpc/boilerplate/internal/web/req"
	"github.com/stretchr/testify/require"
)

func TestValidateCmdItem(t *testing.T) {
	validate := validator.New()
	validate.RegisterStructValidation(ValidateCmdItem, req.CmdItem{})
	configStr := "{\"date\":{\"format\":\"\"}}"
	err := validate.Struct(req.CmdItem{
		Type:   "dateTime",
		Value:  "{}",
		Config: &configStr,
	})
	t.Logf("validate error:%v", err)
	require.Error(t, err)
}
