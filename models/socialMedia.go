package models

import (
	"time"
)

type SocialMedia struct {
	ID             uint   `gorm:"primaryKey" json:"id"`
	Name           string `gorm:"not null" json:"name" form:"name"`
	SocialMediaURL string `gorm:"not null" json:"social_media_url" form:"social_media_url"`
	UserID         uint
	// User           User      `gorm:"foreignKey:ID"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}