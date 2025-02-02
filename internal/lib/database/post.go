package database

import (
	"fmt"

	"github.com/anggunpermata/patreon-clone/configs"
	"github.com/anggunpermata/patreon-clone/internal/models"
	"github.com/labstack/echo/v4"
)

func AddNewPost(c echo.Context, cfg *configs.Config, post models.Post) error {
	if err := cfg.DB.Create(&post).Error; err != nil {
		return err
	}
	return nil
}

func CreateOrUpdatePost(c echo.Context, cfg *configs.Config, postID uint, postData models.Post) error {
	if cfg.DB.Model(&postData).Where("id=?", postID).Updates(&postData).RowsAffected == 0 {
		return fmt.Errorf("failed to update post data with id %v, rows affected = 0", postID)
	}
	return nil
}

func GetAllPostsByUserID(c echo.Context, cfg *configs.Config, userID uint) ([]models.Post, error){
	var posts []models.Post
	if err := cfg.DB.Where("user_id=?", userID).Find(&posts).Error; err != nil {
		return posts, err
	}

	return posts, nil
}

func GetDashboardPostsByUserIDs(c echo.Context, cfg *configs.Config, userIds []models.UserID) ([]models.Post, error) {
	var posts []models.Post

	if err := cfg.DB.Where("user_id IN ?", userIds).Order("created_at DESC").Find(&posts).Error; err != nil {
		return posts, err
	}
	return posts, nil
}

func DeleteOnePostByID(c echo.Context, cfg *configs.Config, postID uint, userID uint) error {
	err := cfg.DB.Where("id = ? AND user_id = ?", postID, userID).Delete(&models.Post{}).Error
	return err
}