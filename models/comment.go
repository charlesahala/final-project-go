package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	Comment_id uint      `gorm:"primaryKey"`
	User_id    []User    `gorm:"foreignKey:User_id" json:"user_id"`
	Photo_id   []Photo   `gorm:"foreignKey:Photo_id" json:"photo_id"`
	Message    string    `gorm:"not null" json:"message" form:"message" valid:"required~message is required"`
	Created_at time.Time `json:"createdAt"`
	Updated_at time.Time `json:"updatedAt"`
}

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(c)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
