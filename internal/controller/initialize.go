package controller

import (

	"github.com/anggunpermata/patreon-clone/configs"
	"github.com/anggunpermata/patreon-clone/middleware/logger"
	"github.com/labstack/echo/v4"
)

type BackendHandler struct {
	cfg         *configs.Config
	logger      logger.Logger
	InitHandler InitHandler
}

func InitiateHandler(cfg *configs.Config, logger logger.Logger) *BackendHandler {
	return &BackendHandler{
		cfg:    cfg,
		logger: logger,
	}
}

type InitHandler interface {
	Healthcheck(c echo.Context)
	RouteUploadFile(c echo.Context)
	UploadFile(c echo.Context)
	UserLogin(c echo.Context)
	UserSignup(c echo.Context)
}
