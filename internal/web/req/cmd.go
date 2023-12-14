package req

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

type DateTimeConfigItem struct {
	Enable *bool   `json:"enable,omitempty"`
	Format *string `json:"format,omitempty" validate:"min=1"`
}

type DateTimeConfig struct {
	Format *string         `json:"format,omitempty" validate:"min=1"`
	Date   *DateTimeConfig `json:"date,omitempty"`
	Time   *DateTimeConfig `json:"time,omitempty"`
}

func ValidateCmdItem(sl validator.StructLevel) {
	cmdItem := sl.Current().Interface().(CmdItem)
	switch cmdItem.Type {
	case "dateTime":
		config := DateTimeConfig{}
		json.Unmarshal([]byte(*cmdItem.Config), &config)
		err := sl.Validator().Struct(config)
		if err != nil {
			sl.ReportValidationErrors("Config", "", err.(validator.ValidationErrors))
		}
	}
}

type CmdCallDto struct {
	Script string `json:"script" validate:"required,min=1"`
	Wd     string `json:"wd" validate:"required,min=1"`
}

type CmdListFolderDto struct {
	Prefix string `json:"prefix" validate:"required,min=1"`
}

type CmdItem struct {
	Type   string  `json:"type" validate:"required,oneof=exec path dateTime"`
	Value  string  `json:"value" validate:"required,min=1"`
	Config *string `json:"config" validate:"min=1"`
}

type CmdAddDto struct {
	Items []CmdItem `json:"items" validate:"required,min=1"`
}

type CmdUpdateDto struct {
	ID    uint32    `json:"id" validate:"required,min=1"`
	Items []CmdItem `json:"items" validate:"min=1"`
	Name  *string   `json:"name" validate:"min=1"`
}

type CmdDeleteDto struct {
	ID uint32 `json:"id" validate:"required,min=1"`
}

type CmdListDto struct {
	Pager
}
