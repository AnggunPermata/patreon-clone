package models

import "gorm.io/gorm"

type SubscribtionInfo struct {
	gorm.Model
	CreatorID            uint    `json:"creator_id"`
	IsFreeFOrAllFollower bool    `json:"is_free_for_all_follower" form:"is_free_for_all_follower"`
	SubscribtionName     string  `json:"subscription_name" form:"subscription_name"`
	IsSubscriptionActive bool    `json:"is_subscription_active" form:"is_subscription_active"`
	Price                float32 `json:"price" form:"price"`
}

type SubscriberInfo struct {
	gorm.Model
	UserID               uint
	SubscriptionID       uint
	IsSubscribing        bool
	BillingStatusPaidOff bool `json:"billing-status-paid-off"`
}
