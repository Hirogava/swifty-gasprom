package game

import (
	"net/http"

	"github.com/Hirogava/swifty-gasprom/backend/internal/config/logger"
	"github.com/Hirogava/swifty-gasprom/backend/internal/repository/postgres"

	"github.com/gin-gonic/gin"
)

func NextMonthMove(c *gin.Context, manager *postgres.Manager) {
	logger.Logger.Debug("Getting next month move", "user_id", c.GetString("userID"), "ip", c.ClientIP())
	newMonth, err := manager.NextMonthMove(c.GetString("userID"))
	switch err {
	case nil:
		logger.Logger.Debug("Next month move retrieved successfully", "user_id", c.GetString("userID"), "ip", c.ClientIP())
		c.JSON(http.StatusOK, gin.H{
			"newMonth": newMonth,
		})
		return
	default:
		logger.Logger.Error("Failed to get next month move", "user_id", c.GetString("userID"), "ip", c.ClientIP(), "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
}
