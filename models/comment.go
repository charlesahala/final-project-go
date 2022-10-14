package models

import "time"

type Comment struct {
	comment_id int `gorm:"primaryKey"`
	user_id    []User
	photo_id   []Photo
	message    string
	created_at time.Time
	updated_at time.Time
}