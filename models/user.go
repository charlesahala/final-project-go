package models

import "time"

type User struct {
	user_id    int    `gorm:"primaryKey"`
	username   string `gorm:"not null" valid:"required~username is required"`
	email      string
	password   string
	age        int
	created_at time.Time
	updated_at time.Time
}
