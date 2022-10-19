package models

import (
	"time"
)

type SocialMedia struct {
	ID             uint   `gorm:"primaryKey" json:"id"`
	Name           string `gorm:"not null" json:"name" form:"name"`
	SocialMediaURL string `gorm:"not null" json:"social_media_url" form:"social_media_url"`
	UserID         uint
	CreatedAt      *time.Time `json:"created_at,omitempty"`
	UpdatedAt      *time.Time `json:"updated_at,omitempty"`
	User           *User       `json:",omitempty"`
}
