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

	if c.Request().Method == "GET" {

		

	} else if c.Request().Method == "POST" {
		var subscriptionInfo models.SubscribtionInfo
		if err := c.Bind(&subscriptionInfo); err != nil {
			return c.JSON(http.StatusBadRequest, "failed to bind subscription info")
		}

		subscriptionInfo.CreatorID = userData["id"].(uint)
		if errCUS := database.CreateOrUpdateSubscription(c, b.cfg.DB, subscriptionInfo); err != nil {
			b.logger.Errorw("failed to create / update subscription", "creator_id", subscriptionInfo.CreatorID, "subscription_name", subscriptionInfo.SubscribtionName, "error", errCUS)
			return c.JSON(http.StatusInternalServerError, "failed to create/update subscription info")
		}

		return c.JSON(201, "successfully create/update new subscription")
	}

	return c.JSON(http.StatusForbidden, "page not found")
}
