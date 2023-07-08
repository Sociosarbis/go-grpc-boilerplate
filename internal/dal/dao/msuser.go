package dao

type MsUser struct {
	ID     string `gorm:"type:varchar(255);primaryKey"`
	UserID uint   `gorm:"column:userId"`
	User   User   `gorm:"foreignKey:userId"`
}

func (*MsUser) TableName() string {
	return "ms_user"
}
