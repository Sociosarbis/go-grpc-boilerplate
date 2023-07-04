package dao

import (
	"time"

	"gorm.io/datatypes"
)

type CommandDataItem struct {
	Type  string         `json:"type"`
	Value map[string]any `json:"value"`
}

type CommandData struct {
	Items []CommandDataItem `json:"items"`
}

type Command struct {
	Common
	Data      datatypes.JSONType[CommandData]
	CreatorId uint32 `gorm:"column:creatorId"`
	Creator   User   `gorm:"foreginKey:creatorId"`
	Users     []User `gorm:"many2many:user_commands_command"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
