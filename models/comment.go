package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	ID     uint `gorm:"primaryKey" json:"id"`
	UserID uint
	// User      User `gorm:"foreignKey:ID"`
	PhotoID uint `json:"photo_id"`
	// Photo     Photo     `gorm:"foreignKey:ID"`
	Message   string    `gorm:"not null" json:"message" form:"message"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
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
