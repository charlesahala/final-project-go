package models

import "time"

type Photo struct {
	photo_id   int `gorm:"primaryKey"`
	title      string
	caption    string
	photo_url  string
	user_id    []User
	created_at time.Time
	updated_at time.Time
}
