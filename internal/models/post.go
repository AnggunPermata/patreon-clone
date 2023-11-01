package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	UserID           uint   `json:"user_id"`
	Text             string `json:"text"`
	NeedSubscribtion bool   `json:"need_subscription"`
	SubscriptionID   uint   `json:"subscription_id"`
}
