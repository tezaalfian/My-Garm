package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// Photo represents a photo
type Photo struct {
	GormModel
	Title    string `gorm:"not null" json:"title" form:"title" valid:"required"`
	Caption  string `gorm:"null" json:"caption"`
	PhotoUrl string `gorm:"not null" json:"photo_url" form:"photo_url" valid:"required,url"`
	UserID   uint
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)
	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return
}
