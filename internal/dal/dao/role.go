package dao

type Role struct {
	Common
	Name string `gorm:"type:varchar(255);"`
}

func (*Role) TableName() string {
	return "role"
}
