package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	Photo_id   uint      `gorm:"primaryKey" json:"photo_id"`
	Title      string    `gorm:"not null" json:"title" form:"title" valid:"required~title is required"`
	Caption    string    `json:"caption" form:"caption"`
	Photo_url  string    `gorm:"not null" json:"photo_url" form:"photo_url" valid:"required~photo url is required"`
	User_id    []User    `gorm:"foreignKey:User_id" json:"user_id"`
	Created_at time.Time `json:"createdAt"`
	Updated_at time.Time `json:"updatedAt"`
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
