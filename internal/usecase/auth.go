package usecase

import (
	"fmt"

	"github.com/anggunpermata/patreon-clone/auth"
	"github.com/anggunpermata/patreon-clone/configs"
	"github.com/anggunpermata/patreon-clone/internal/lib/database"
	"github.com/labstack/echo/v4"
)

func CheckingAuthorization(c echo.Context, cfg *configs.Config, cookieName string) (map[string]interface{}, error) {
	authCookie, err := c.Cookie(cookieName)
	if err != nil {
		return nil, err
	}

	tokenStr := authCookie.Value
	role, username, authorized := auth.IdentifyUserUsingJWTToken(c, cfg, tokenStr)
	if (role == "admin" || role == "user") && authorized {
		user, err := database.GetOneUserByUsername(c, cfg.DB, username)
		if err == nil {
			showUserData := map[string]interface{}{
				"id":       user.ID,
				"fullName": user.FullName,
				"username": user.Username,
				"email":    user.Email,
				"role":     user.Role,
				"token":    user.Token,
			}
			return showUserData, nil
		}
	}

	return nil, fmt.Errorf("unauthorized")
}
