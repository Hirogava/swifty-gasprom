package http

import (
	"time"

	"github.com/Hirogava/swifty-gasprom/backend/internal/config/logger"
	"github.com/Hirogava/swifty-gasprom/backend/internal/handler/analytics"
	"github.com/Hirogava/swifty-gasprom/backend/internal/handler/auth"
	"github.com/Hirogava/swifty-gasprom/backend/internal/handler/bank"
	"github.com/Hirogava/swifty-gasprom/backend/internal/handler/game"
	"github.com/Hirogava/swifty-gasprom/backend/internal/repository/postgres"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func CreateRouter(manager *postgres.Manager) *gin.Engine {
	logger.Logger.Debug("Creating HTTP router")

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

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
