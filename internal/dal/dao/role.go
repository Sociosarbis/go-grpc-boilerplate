package dao

type Role struct {
	Common
	Name string
}

func (*Role) TableName() string {
	return "role"
}
