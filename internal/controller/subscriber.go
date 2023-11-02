package controller

import (
	"net/http"

	"github.com/anggunpermata/patreon-clone/internal/lib/database"
	"github.com/anggunpermata/patreon-clone/internal/models"
	"github.com/anggunpermata/patreon-clone/internal/usecase"
	"github.com/labstack/echo/v4"
)

func (b *BackendHandler) CreateANewSubscription(c echo.Context) error {
	userData, err := usecase.CheckingAuthorization(c, b.cfg, "Authorization")
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if c.Request().Method == "POST" {
		var subscriptionInfo models.SubscribtionInfo
		if err := c.Bind(&subscriptionInfo); err != nil {
			return c.JSON(http.StatusBadRequest, "failed to bind subscription info")
		}

		subscriptionInfo.CreatorID = userData["id"].(uint)
		
	}
}
