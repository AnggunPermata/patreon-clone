package usecase

import (
	"net/http"

	"github.com/anggunpermata/patreon-clone/configs"
	"github.com/anggunpermata/patreon-clone/internal/lib/database"
	"github.com/anggunpermata/patreon-clone/internal/models"
	"github.com/labstack/echo/v4"
)

func CreateAPost(c echo.Context) {

}

func GetAllPostsByUserID(c echo.Context, cfg *configs.Config, userID uint) ([]models.Post, error) {
	allPost, err := database.GetAllPostsByUserID(c, cfg, userID)
	if err != nil {
		return allPost, c.JSON(http.StatusBadRequest, map[string]string{"status": "failed to get all posts", "error_message": err.Error()})
	}

	return allPost, nil
}
