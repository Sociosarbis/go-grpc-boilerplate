package dao

type User struct {
	Common
	Name   string
	Groups []Group `gorm:"many2many:group_users_user;"`
}

func (*User) TableName() string {
	return "user"
}
