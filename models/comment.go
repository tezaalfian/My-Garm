package models

type Comment struct {
	GormModel
	Message string `gorm:"not null" json:"message" form:"message" valid:"required"`
	PhotoID uint   `gorm:"not null" json:"photo_id" form:"photo_id" valid:"required"`
	UserID  uint
}
