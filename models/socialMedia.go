package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	Social_media_id  uint   `gorm:"primaryKey"`
	Name             string `gorm:"not null" json:"name" form:"name" valid:"required~name is required"`
	Social_media_url string `gorm:"not null" json:"social_media_url" form:"social_media_url" valid:"required~name is required"`
	User_id          []User `gorm:"foreignKey:User_id" json:"user_id"`
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
