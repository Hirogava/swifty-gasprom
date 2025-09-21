package analytics

import (
	"github.com/Hirogava/swifty-gasprom/backend/internal/repository/postgres"
	
	"github.com/gin-gonic/gin"
)

func InitAnalyticsHandlers(r *gin.Engine, manager *postgres.Manager) {}