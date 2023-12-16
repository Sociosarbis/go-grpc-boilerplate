package cmd

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
	"github.com/sociosarbis/grpc/boilerplate/internal/web/req"
)

func validateJsonField[T any](sl validator.StructLevel, jsonValue string) {
	var value T
	json.Unmarshal([]byte(jsonValue), &value)
	err := sl.Validator().Struct(value)
	if err != nil {
		sl.ReportValidationErrors("", "", err.(validator.ValidationErrors))
	}
}

func ValidateCmdItem(sl validator.StructLevel) {
	cmdItem := sl.Current().Interface().(req.CmdItem)
	if cmdItem.Type != "exec" {
		temp := map[string]any{}
		err := json.Unmarshal([]byte(cmdItem.Value), &temp)
		if err != nil {
			sl.ReportError(cmdItem.Value, "Value", "", "json", "")
		}
	}
	switch cmdItem.Type {
	case "path":
		validateJsonField[DateTimeValue](sl, cmdItem.Value)
	case "dateTime":
		validateJsonField[DateTimeConfig](sl, *cmdItem.Config)
		validateJsonField[DateTimeValue](sl, cmdItem.Value)
	}
}
