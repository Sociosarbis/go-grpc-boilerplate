package dao

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type CommandDataItem struct {
	Type  string
	Value string
}

type CommandData struct {
	Items []CommandDataItem
}

type Command struct {
	gorm.Model
	Data      datatypes.JSONType[CommandData]
	CreatorId uint   `gorm:"column:creatorId"`
	Creator   User   `gorm:"foreginKey:creatorId"`
	Users     []User `gorm:"many2many:user_commands_command"`
}
