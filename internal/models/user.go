package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username    string `json:"username" form:"username"`
	Email       string `json:"email" form:"email"`
	Password    string `json:"password" form:"password"`
	FullName    string `json:"full-name" form:"full-name"`
	Gender      string `json:"gender" form:"gender"`
	Role        string `json:"role" form:"role"`
	ProfilePage string `json:"profile_page"`
	AvatarUrl   string `json:"avatar_url"`
	Token       string
}
