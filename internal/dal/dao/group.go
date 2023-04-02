package dao

type Group struct {
	Common
	Name     string
	ParentID uint32  `gorm:"column:parentId"`
	Children []Group `gorm:"foreignKey:parentId"`
	Roles    []Role  `gorm:"many2many:group_roles_role"`
}

func (*Group) TableName() string {
	return "group"
}
