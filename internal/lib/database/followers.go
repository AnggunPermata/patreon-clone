package database

import (
	"fmt"

	"github.com/anggunpermata/patreon-clone/configs"
	"github.com/anggunpermata/patreon-clone/internal/models"
	"github.com/labstack/echo/v4"
)

func AddFollowingInfoByUserID(c echo.Context, cfg *configs.Config, userID uint, followedID uint) error {
	var followingData = models.Followers{
		UserID:         userID,
		FollowedUserID: followedID,
		IsFollowing:    true,
	}

	if cfg.DB.Model(&followingData).Where("user_id=? AND followed_user_id=?", userID, followedID).Updates(&followingData).RowsAffected == 0 {
		if err := cfg.DB.Create(&followingData).Error; err != nil {
			return err
		}
	}

	return nil
}

func GetOneFollowerByID(c echo.Context, cfg *configs.Config, userID uint, followedID uint) (models.Followers, error) {
	var followerData models.Followers
	if err := cfg.DB.Where("user_id=? AND followed_user_id=?", userID, followedID).First(&followerData).Error; err != nil {
		return followerData, err
	}

	if followerData.IsFollowing {
		return followerData, nil
	} else {
		return followerData, fmt.Errorf("error when trying to get one follower by id, isFollowing != true. userID = %v followedID = %v", userID, followedID)
	}
}

func GetAllFollowerByUserID(c echo.Context, cfg *configs.Config, userID uint) ([]models.Followers, error) {
	var allFollower []models.Followers
	if err := cfg.DB.Where("followed_user_id=?", userID).Find(&allFollower).Error; err != nil {
		return allFollower, err
	}

	return allFollower, nil
}

func GetAllFollowerID(c echo.Context, cfg *configs.Config, userID uint) ([]models.UserID, error) {
	var followersID []models.UserID
	sql := fmt.Sprintf("select user_id from followers where followed_user_id = %d", userID)
	if err := cfg.DB.Raw(sql).Scan(&followersID).Error; err != nil {
		return followersID, err
	}

	return followersID, nil
}

func GetAllFollowingID(c echo.Context, cfg *configs.Config, userID uint) ([]models.UserID, error) {
	var followingID []models.UserID
	sql := fmt.Sprintf("select followed_user_id from followers where user_id = %d", userID)
	if err := cfg.DB.Raw(sql).Scan(&followingID).Error; err != nil {
		return followingID, err
	}

	return followingID, nil
}
