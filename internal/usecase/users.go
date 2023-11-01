package usecase

import (
	"fmt"

	"github.com/anggunpermata/patreon-clone/auth"
	"github.com/anggunpermata/patreon-clone/configs"
	"github.com/anggunpermata/patreon-clone/internal/lib/database"
	"github.com/anggunpermata/patreon-clone/internal/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func UserLoginWithEmail(c echo.Context, cfg *configs.Config, DB *gorm.DB, email string, password string) (models.User, error) {
	user, err := database.GetUserByEmailAndPassword(c, cfg, DB, email, password)
	if err != nil {
		return user, fmt.Errorf("user failed to login user with email=%s, error_message=%v", email, err)
	}

	user.Token, err = auth.CreateToken(c, cfg.SecretJWT, int(user.ID), user.Role, user.Username, 12)
	if err != nil {
		return user, fmt.Errorf("failed to create token for user with email=%s, error_message=%v", email, err)
	}

	return user, nil
}

func UserSignUp(c echo.Context, cfg *configs.Config, userData models.User) error {
	return nil
}
