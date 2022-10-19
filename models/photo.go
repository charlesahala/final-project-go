package models

import "time"


type Photo struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Title     string `gorm:"not null" json:"title" form:"title"`
	Caption   string `json:"caption" form:"caption"`
	PhotoURL  string `gorm:"not null" json:"photo_url" form:"photo_url"`
	UserID    uint
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	User      User
	Comments  []Comment `json:"-"`
}