package models

type SocialMedia struct {
	social_media_id int `gorm:"primaryKey"`
	name string
	social_media_url string
	user_id []User
}