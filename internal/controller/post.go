package controller

import (
	"net/http"
	"strconv"

	"github.com/anggunpermata/patreon-clone/internal/lib/database"
	"github.com/anggunpermata/patreon-clone/internal/models"
	"github.com/anggunpermata/patreon-clone/internal/usecase"
	"github.com/labstack/echo/v4"
)

type ReturnPostData struct {
	PostInfo []models.Post
}

func (b *BackendHandler) CreateAPost(c echo.Context) error {
	if c.Request().Method == "GET" {
		userMapData, err := usecase.CheckingAuthorization(c, b.cfg, "Authorization")
		if err != nil {
			return c.Render(200, "login.html", nil)
		}

		return c.Render(200, "posting.html", userMapData)
	}

	if c.Request().Method == "POST" {
		userMap, err := usecase.CheckingAuthorization(c, b.cfg, "Authorization")
		if err != nil {
			return c.Render(200, "login.html", nil)
		}

		postData := models.Post{}
		postData.UserID = userMap["id"].(uint)
		postData.Text = c.FormValue("create_post")

		if err := database.AddNewPost(c, b.cfg, postData); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"status": "failed to create a new post", "error_message": err.Error()})
		}

		return c.JSON(200, map[string]interface{}{"status": "success"})
	}

	return c.JSON(404, map[string]interface{}{"status": "page not found"})
}

func (b *BackendHandler) GetAllPostsByUserID(c echo.Context) error {
	targetedUserId, _ := strconv.Atoi(c.Param("targeted_user_id"))
	allPosts, err := usecase.GetAllPostsByUserID(c, b.cfg, uint(targetedUserId))
	if err != nil {
		return err
	}

	return c.Render(200, "all-post.html", ReturnPostData{
		PostInfo: allPosts,
	})
}
