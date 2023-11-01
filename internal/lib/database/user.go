package database

import (
	"fmt"

	"github.com/anggunpermata/patreon-clone/configs"
	"github.com/anggunpermata/patreon-clone/internal/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateOrUpdateOneUser(c echo.Context, DB *gorm.DB, user models.User) (*models.User, error) {
	// create if record is not exists
	if DB.Model(&user).Where("username=?", user.Username).Updates(&user).RowsAffected == 0 {
		DB.Create(&user)
	}

	return &user, nil
}

// Get only one User by username
func GetOneUserByUsername(c echo.Context, DB *gorm.DB, username string) (models.User, error) {
	var user models.User

	if err := DB.Where("username=?", username).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

// Get only one User by id
func GetOneUserByUserID(c echo.Context, DB *gorm.DB, id uint) (models.User, error) {
	var user models.User

	if err := DB.Where("id=?", id).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func GetUserByEmailAndPassword(c echo.Context, cfg *configs.Config, DB *gorm.DB, email string, password string) (models.User, error) {
	var user models.User
	var err error

	if err = DB.Where("email=? AND password=?", email, password).First(&user).Error; err != nil {
		return user, fmt.Errorf("failed to sign user with email=%s, error_message=%v", email, err)
	}

	return user, nil
}

func GetMultipleUserByUserIDs(c echo.Context, cfg *configs.Config, userIds []models.UserID) ([]models.User, error) {
	var users []models.User

	if err := cfg.DB.Where("id IN ?", userIds).Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}

func ResetToken(c echo.Context, cfg *configs.Config, user models.User) error {
	user.Token = ""
	if err := cfg.DB.Model(&user).Updates(&user).Error; err != nil {
		return err
	}

	return nil
}
