package models

type Photo struct {
	GormModel
	Title    string `gorm:"not null" json:"title" form:"title" valid:"required"`
	Caption  string `gorm:"null" json:"caption"`
	PhotoUrl string `gorm:"not null" json:"photo_url" form:"photo_url" valid:"required"`
	UserID   uint
}
