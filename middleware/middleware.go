package middleware

import (
	"github.com/anggunpermata/patreon-clone/configs"
	"github.com/anggunpermata/patreon-clone/middleware/logger"
)

// Middleware manager
type MiddlewareManager struct {
	cfg *configs.Config
	logger logger.Logger
}

// To create a new MiddlewareManager
func NewMiddlewareManager(cfg *configs.Config, logger logger.Logger) *MiddlewareManager {
	return &MiddlewareManager{
		cfg:    cfg,
		logger: logger,
	}
}