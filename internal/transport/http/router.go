package http

import (
	"github.com/Hirogava/swifty-gasprom/backend/internal/handler/analytics"
	"github.com/Hirogava/swifty-gasprom/backend/internal/handler/auth"
	"github.com/Hirogava/swifty-gasprom/backend/internal/handler/bank"
	"github.com/Hirogava/swifty-gasprom/backend/internal/handler/game"
	"github.com/Hirogava/swifty-gasprom/backend/internal/repository/postgres"
	"github.com/gin-gonic/gin"
)

func CreateRouter(manager *postgres.Manager) *gin.Engine {
	r := gin.Default()

	game.InitGameHandlers(r, manager)
	bank.InitBankHandlers(r, manager)
	auth.InitAuthHandlers(r, manager)
	analytics.InitAnalyticsHandlers(r, manager)

	return r
}