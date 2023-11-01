package models

import "gorm.io/gorm"

type SubscribtionInfo struct {
	gorm.Model
	CreatorID            string  `json:"creator-id"`
	SubscribtionType     string  `json:"type"`
	SubscribtionName     string  `json:"name"`
	IsSubscriptionActive bool    `json:"is_subscription_active"`
	Price                float32 `json:"price"`
	BillingStatusPaidOff bool    `json:"billing-status-paid-off"`
}

type SubscriberInfo struct {
	gorm.Model
	UserID         uint
	SubscriptionID uint
	IsSubscribing  bool
}
