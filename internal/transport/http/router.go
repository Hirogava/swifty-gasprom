package http

import (
	"github.com/Hirogava/swifty-gasprom/backend/internal/config/logger"
	"github.com/Hirogava/swifty-gasprom/backend/internal/handler/analytics"
	"github.com/Hirogava/swifty-gasprom/backend/internal/handler/auth"
	"github.com/Hirogava/swifty-gasprom/backend/internal/handler/bank"
	"github.com/Hirogava/swifty-gasprom/backend/internal/handler/game"
	"github.com/Hirogava/swifty-gasprom/backend/internal/repository/postgres"
	"github.com/gin-gonic/gin"
)

func CreateRouter(manager *postgres.Manager) *gin.Engine {
	logger.Logger.Debug("Creating HTTP router")

	r := gin.Default()

	logger.Logger.Debug("Registering game handlers")
	game.InitGameHandlers(r, manager)

	logger.Logger.Debug("Registering bank handlers")
	bank.InitBankHandlers(r, manager)

	logger.Logger.Debug("Registering auth handlers")
	auth.InitAuthHandlers(r, manager)

	logger.Logger.Debug("Registering analytics handlers")
	analytics.InitAnalyticsHandlers(r, manager)

	logger.Logger.Info("HTTP router created successfully")
	return r
}
