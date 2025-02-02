package controller

import (
	"net/http"
	"strings"

	"github.com/anggunpermata/patreon-clone/auth"
	"github.com/anggunpermata/patreon-clone/internal/cookies"
	"github.com/anggunpermata/patreon-clone/internal/lib/database"
	"github.com/anggunpermata/patreon-clone/internal/models"
	"github.com/anggunpermata/patreon-clone/internal/usecase"
	"github.com/labstack/echo/v4"
)

func AuthorizedUser(c echo.Context) bool {
	_, role := auth.ExtractTokenUserId(c)
	if role != "user" && role != "admin" {
		return false
	}
	return true
}

func (b *BackendHandler) UserSignup(c echo.Context) error {
	if c.Request().Method == "GET" {
		role := "user"
		return c.Render(200, "signup.html", role)
	}
	userData := models.User{}
	if err := c.Bind(&userData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"status": "failed to bind user data"})
	}

	if len(userData.Username) < 4 || len(userData.Email) < 4 || !strings.Contains(userData.Email, ".com") || len(userData.Password) < 6 || len(userData.Gender) > 1 || (userData.Gender != "F" && userData.Gender != "M") {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Please follow the rules to Sign up:",
			"rules-1": "Make sure the username has more than 3 characters.",
			"rules-2": "Make sure the email has more than 3 characters, and it is a real email.",
			"rules-3": "Make sure the Password has more than 5 characters.",
			"rules-4": "Gender only have one character, which is F for female or M for male",
		})

	}

	user, err := database.GetOneUserByUsername(c, b.cfg.DB, userData.Username)
	if err == nil && len(user.Username) > 0 {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Cannot Signup, username has already taken",
		})
	}

	signup, err := database.CreateOrUpdateOneUser(c, b.cfg.DB, userData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message":       "Cannot Signup",
			"error_message": err.Error(),
		})
	}

	showUserData := map[string]interface{}{
		"ID":        signup.ID,
		"Name":      signup.FullName,
		"Email":     signup.Email,
		"Full Name": signup.FullName,
		"Gender":    signup.Gender,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Succesfully create a new account",
		"data":    showUserData,
	})
}

func (b *BackendHandler) UserLogin(c echo.Context) error {
	if c.Request().Method == "GET" {
		authCookie, err := c.Cookie("Authorization")
		if err != nil {
			return c.Render(200, "login.html", nil)
		}

		tokenStr := authCookie.Value
		role, username, authorized := auth.IdentifyUserUsingJWTToken(c, b.cfg, tokenStr)
		if (role == "admin" || role == "user") && authorized {
			user, err := database.GetOneUserByUsername(c, b.cfg.DB, username)
			if err == nil {
				posts, err := usecase.GetAllPostsByUserID(c, b.cfg, user.ID)
				if err != nil {
					return c.JSON(http.StatusBadRequest, "failed to load posts")
				}
				return c.Render(200, "dashboard.html", ReturnUserData{
					User:     user,
					PostData: posts,
				})
			}
		}

		return c.Render(200, "login.html", nil)
	}

	inputData := models.User{}
	if err := c.Bind(&inputData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"status": "failed to bind user data"})
	}

	loginData := models.User{
		Email:    inputData.Email,
		Password: inputData.Password,
	}

	user, err := usecase.UserLoginWithEmail(c, b.cfg, b.cfg.DB, loginData.Email, loginData.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message":       "please check your email and password again.",
			"error_message": err.Error(),
		})
	}

	posts, err := usecase.GetAllPostsByUserID(c, b.cfg, user.ID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "failed to load posts")
	}

	cookies.WriteCookie(c, "Authorization", user.Token, "/", 2400)

	return c.Render(200, "dashboard.html", ReturnUserData{
		User:     user,
		PostData: posts,
	})

}

func (b *BackendHandler) UserLogout(c echo.Context) error {
	if c.Request().Method == "POST" || c.Request().Method == "GET" {
		authCookie, err := c.Cookie("Authorization")
		if err != nil {
			return c.JSON(http.StatusBadRequest, "error fetching auth information")
		}

		tokenStr := authCookie.Value
		role, username, authorized := auth.IdentifyUserUsingJWTToken(c, b.cfg, tokenStr)
		if (role == "admin" || role == "user") && authorized {
			user, err := database.GetOneUserByUsername(c, b.cfg.DB, username)
			if err == nil {
				// reset token
				if err := database.ResetToken(c, b.cfg, user); err != nil {
					b.logger.Errorf("failed to reset user id %v token with error : %v", user.ID, err.Error())
					return c.JSON(http.StatusInternalServerError, "error occured when trying to log account off : failed to reset token")
				}

				// overwrite cookies
				cookies.WriteCookie(c, "Authorization", "nil", "/", 1)
				return c.Render(200, "logout.html", nil)
			}
		}

		return c.JSON(http.StatusBadRequest, "error occured when trying to log account off : unauthorized user")
	}

	return c.JSON(http.StatusNotFound, "path not found")
}
