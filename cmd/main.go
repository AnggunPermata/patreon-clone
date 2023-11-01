package main

import (
	"fmt"

	"github.com/anggunpermata/patreon-clone/configs"
	"github.com/anggunpermata/patreon-clone/internal/controller"
	"github.com/anggunpermata/patreon-clone/internal/routes"
	"github.com/anggunpermata/patreon-clone/middleware/logger"
	"github.com/anggunpermata/patreon-clone/template"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// load config
	cfg := configs.LoadConfig()

	// setup logger
	appLogger := logger.NewLoggerConfig(cfg)
	appLogger.InitLogger()

	// Setup MiddlewareManager
	// middlewareManager := appMiddleware.NewMiddlewareManager(cfg, appLogger)

	// setup route engine
	e := echo.New()
	e.Use(middleware.CORS())
	// e.Use(middlewareManager.InboundLog)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/static", "assets")

	// render template
	template.TemplateRenderer(e)

	// setup controller
	handlerUsecase := controller.InitiateHandler(cfg, appLogger)
	routes.NewRoutes(e, handlerUsecase)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.ServerPort)))
}
