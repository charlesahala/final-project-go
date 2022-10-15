package models

import (
	"final-project-go/helpers"
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	User_id    uint      `gorm:"primaryKey" json:"user_id"`
	Username   string    `gorm:"not null;uniqueIndex" json:"username" form:"username" valid:"required~username is required"`
	Email      string    `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~email is required,email~invalid email format"`
	Password   string    `gorm:"not null" json:"password" form:"password" valid:"required~password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Age        int       `gorm:"not null,where:age > 7" json:"age" form:"age" valid:"required~age is required and must be 8+ years old"`
	Created_at time.Time `json:"createdAt"`
	Updated_at time.Time `json:"updateedAt"`
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
