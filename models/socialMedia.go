package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// SocialMedia represents a social media account
type SocialMedia struct {
	GormModel
	Name           string `gorm:"not null" json:"name" form:"name" valid:"required"`
	SocialMediaUrl string `gorm:"not null" json:"social_media_url" form:"social_media_url" valid:"required,url"`
	UserID         uint
}

func (s *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(s)
	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return
}
