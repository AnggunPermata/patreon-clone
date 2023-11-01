package controller

import (
	"net/http"
	"strconv"

	"github.com/anggunpermata/patreon-clone/internal/lib/database"
	"github.com/anggunpermata/patreon-clone/internal/models"
	"github.com/anggunpermata/patreon-clone/internal/usecase"
	"github.com/labstack/echo/v4"
)

type ReturnUserData struct {
	User     models.User
	PostData []models.Post
}

func (b *BackendHandler) UserProfiles(c echo.Context) error {
	if c.Request().Method == "GET" {
		renderPage := "other-profile.html"
		targetUserId, _ := strconv.Atoi(c.Param("targeted_user_id"))
		userData, err := usecase.CheckingAuthorization(c, b.cfg, "Authorization")
		if err == nil {
			if userData["id"].(uint) == uint(targetUserId) {
				renderPage = "user-profile.html"
			}
		}
		user, err := database.GetOneUserByUserID(c, b.cfg.DB, uint(targetUserId))
		if err != nil {
			return c.JSON(http.StatusBadRequest, "user not found")
		}

		posts, err := usecase.GetAllPostsByUserID(c, b.cfg, uint(targetUserId))
		if err != nil {
			return c.JSON(http.StatusBadRequest, "failed to load posts")
		}

		return c.Render(200, renderPage, ReturnUserData{
			User:     user,
			PostData: posts,
		})
	}

	return nil
}

func (b *BackendHandler) ShowUserDashboard(c echo.Context) error {
	if c.Request().Method == "GET" {
		renderPage := "dashboard.html"
		userData, err := usecase.CheckingAuthorization(c, b.cfg, "Authorization")
		if err != nil {
			renderPage = "login.html"
			return c.Render(200, renderPage, nil)
		}

		// Get current user by id
		currentUser, err := database.GetOneUserByUserID(c, b.cfg.DB, userData["id"].(uint))
		if err != nil {
			return c.Render(200, "signup.html", nil)
		}
		// Get Following id
		followingIDs, err := database.GetAllFollowingID(c, b.cfg, currentUser.ID)
		if err != nil {
			return c.JSON(400, err.Error())
		}
		followingIDs = append(followingIDs, models.UserID(currentUser.ID))
		// Placed here after exist : get subscription id by current user id
		// Get All post by user id and following user id (and subscription id after implemented)
		posts, _ := database.GetDashboardPostsByUserIDs(c, b.cfg, followingIDs)

		return c.Render(200, renderPage, ReturnUserData{
			User: currentUser,
			PostData: posts,
		})
		
	}

	return nil
}
