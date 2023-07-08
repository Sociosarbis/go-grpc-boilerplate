package dao

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name         string  `gorm:"type:varchar(255);"`
	PasswordHash string  `gorm:"type:varchar(255);"`
	Email        string  `gorm:"type:text;"`
	Groups       []Group `gorm:"many2many:group_users_user;"`
}

func (*User) TableName() string {
	return "user"
}
