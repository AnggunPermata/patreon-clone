package models

import "gorm.io/gorm"

type CreatorsMediaActivities struct {
	gorm.Model
	FileName                string `json:"file_name"`
	FileExtension           string `json:"file_extension"`
	FileType                string `json:"file_type"`
	FilePath                string `json:"file_path"`
	SenderID                string `json:"sender_id"`
	SenderEmail             string `json:"sender_email"`
	SpecialSubscriberStatus bool   `json:"special_subscriber_status"`
	SubscriberID            string `json:"subscriber_id"`
}
