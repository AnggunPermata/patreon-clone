package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/anggunpermata/patreon-clone/internal/lib/database"
	"github.com/anggunpermata/patreon-clone/internal/models"
	"github.com/anggunpermata/patreon-clone/internal/usecase"
	"github.com/labstack/echo/v4"
)

type ReturnData struct {
	UserData []models.User
}

func (b *BackendHandler) FollowAUser(c echo.Context) error {
	if c.Request().Method == "POST" || c.Request().Method == "GET" {
		userData, err := usecase.CheckingAuthorization(c, b.cfg, "Authorization")
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		targetedUserId, _ := strconv.Atoi(c.Param("targeted_user_id"))
		targetUser, err := database.GetOneUserByUserID(c, b.cfg.DB, uint(targetedUserId))
		if err != nil {
			return c.JSON(http.StatusBadRequest, fmt.Sprintf("failed to follow the targeted user with error : %s", err.Error()))
		}

		if userData["id"].(uint) == targetUser.ID {
			return c.JSON(http.StatusBadRequest, "failed to follow the targeted user with error : not allowed to follow the same user id")
		}

		followerData, err := database.GetOneFollowerByID(c, b.cfg, userData["id"].(uint), targetUser.ID)
		if err == nil && followerData.IsFollowing {
			return c.JSON(http.StatusBadRequest, "failed to follow the targeted user with error : user already followed")
		}

		// follow the user
		if err := database.AddFollowingInfoByUserID(c, b.cfg, userData["id"].(uint), targetUser.ID); err != nil {
			return c.JSON(http.StatusBadRequest, fmt.Sprintf("failed to follow the targeted user with error : %s", err.Error()))
		}

		return c.JSON(200, fmt.Sprintf("successfully followed a user : %s", targetUser.Username))
	}

	return c.JSON(http.StatusFound, "path not found")
}

func (b *BackendHandler) GetAllFollower(c echo.Context) error {
	if c.Request().Method == "GET" {
		targetedUserId, _ := strconv.Atoi(c.Param("targeted_user_id"))
		allFollowersID, err := database.GetAllFollowerID(c, b.cfg, uint(targetedUserId))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		allUsers, err := database.GetMultipleUserByUserIDs(c, b.cfg, allFollowersID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.Render(200, "show-followers.html", ReturnData{
			UserData: allUsers,
		})
	}

	return c.JSON(http.StatusNotFound, "no path provided")
}

func (b *BackendHandler) GetAllFollowing(c echo.Context) error {
	if c.Request().Method == "GET" {
		targetedUserId, _ := strconv.Atoi(c.Param("targeted_user_id"))
		allFollowingID, err := database.GetAllFollowingID(c, b.cfg, uint(targetedUserId))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		allUsers, err := database.GetMultipleUserByUserIDs(c, b.cfg, allFollowingID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.Render(200, "show-following.html", ReturnData{
			UserData: allUsers,
		})
	}

	return c.JSON(http.StatusNotFound, "no path provided")
}

func (b *BackendHandler) UnfollowAUser(c echo.Context) error {
	if c.Request().Method == "POST" || c.Request().Method == "GET" {
		userData, err := usecase.CheckingAuthorization(c, b.cfg, "Authorization")
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		targetedUserId, _ := strconv.Atoi(c.Param("targeted_user_id"))
		targetUser, err := database.GetOneUserByUserID(c, b.cfg.DB, uint(targetedUserId))
		if err != nil {
			return c.JSON(http.StatusBadRequest, fmt.Sprintf("failed to unfollow the targeted user with error : %s", err.Error()))
		}

		if userData["id"].(uint) == targetUser.ID {
			return c.JSON(http.StatusBadRequest, "failed to unfollow the targeted user with error : not allowed to follow the same user id")
		}

		followerData, err := database.GetOneFollowerByID(c, b.cfg, userData["id"].(uint), targetUser.ID)
		if err != nil && followerData.IsFollowing {
			return c.JSON(http.StatusBadRequest, "failed to unfollow the targeted user with error : need to follow first")
		}

		// follow the user
		if err := database.DeleteOneFollowingByUserID(c, b.cfg, userData["id"].(uint), targetUser.ID); err != nil {
			return c.JSON(http.StatusBadRequest, fmt.Sprintf("failed to unfollow the targeted user with error : %s", err.Error()))
		}

		return c.JSON(200, fmt.Sprintf("successfully unfollowed a user : %s", targetUser.Username))
	}

	return c.JSON(http.StatusFound, "path not found")
}