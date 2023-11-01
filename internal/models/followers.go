package models

import "gorm.io/gorm"


type UserID uint

type Followers struct {
	gorm.Model
	UserID         uint `json:"user_id"`
	FollowedUserID uint `json:"followed_user_id"`
	IsFollowing    bool `json:"is_following"`
}
