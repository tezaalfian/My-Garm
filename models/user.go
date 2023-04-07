package models

import (
	"MyGarm/helpers"
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// user represents a user
type User struct {
	GormModel
	Username string `gorm:"not null;uniqueIndex" json:"username" form:"username" valid:"required"`
	Email    string `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required,email"`
	Password string `gorm:"not null" json:"password" form:"password" valid:"required,minstringlength(6)"`
	Age      int    `gorm:"not null" json:"age" form:"age" valid:"required,numeric,range(8|300)"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)
	if errCreate != nil {
		err = errCreate
		return
	}
	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}
