package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Post
	UserID           uint   `json:"user_id"`
	Text             string `json:"text"`
}
