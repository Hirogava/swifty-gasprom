package game

import (
	"net/http"

	"github.com/Hirogava/swifty-gasprom/backend/internal/config/logger"
	"github.com/Hirogava/swifty-gasprom/backend/internal/repository/postgres"
	
	"github.com/gin-gonic/gin"
)

func GetRandomEvent(c *gin.Context, manager *postgres.Manager) {
	logger.Logger.Debug("Getting random event", "user_id", c.GetString("userID"), "ip", c.ClientIP())
	event, err := manager.GetRandomEvent()
	switch err {
	case nil:
		logger.Logger.Debug("Random event retrieved successfully", "user_id", c.GetString("userID"), "ip", c.ClientIP())
		c.JSON(http.StatusOK, gin.H{
			"event": event,
		})
		return
	default:
		logger.Logger.Error("Failed to get random event", "user_id", c.GetString("userID"), "ip", c.ClientIP(), "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
}