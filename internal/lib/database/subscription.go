package database

import (
	"github.com/anggunpermata/patreon-clone/internal/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateOneSubscription(c echo.Context, DB *gorm.DB, subscriptionData models.SubscribtionInfo) error {
	if err := DB.Create(&subscriptionData).Error; err != nil {
		return err
	}

	return nil
}

func CreateOrUpdateSubscription(c echo.Context, DB *gorm.DB, subscriptionData models.SubscribtionInfo) error {
	var err error
	// create if record is not exists
	if DB.Model(&subscriptionData).Where("id=?", subscriptionData.ID).Updates(&subscriptionData).RowsAffected == 0 {
		err = DB.Create(&subscriptionData).Error
	}

	return err
}

func GetOneSubscriptionByID(c echo.Context, DB *gorm.DB, subscriptionID uint) (models.SubscriberInfo, error) {
	var subscription models.SubscriberInfo

	if err := DB.Where("id=?", subscriptionID).First(&subscription).Error; err != nil {
		return subscription, err
	}
	return subscription, nil
}