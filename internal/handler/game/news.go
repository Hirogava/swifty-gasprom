package game

import (
	"net/http"
	"strconv"

	"github.com/Hirogava/swifty-gasprom/backend/internal/config/logger"
	dbErrors "github.com/Hirogava/swifty-gasprom/backend/internal/errors/db"
	"github.com/Hirogava/swifty-gasprom/backend/internal/repository/postgres"

	"github.com/gin-gonic/gin"
)

func GetPlayerNews(c *gin.Context, manager *postgres.Manager) {
	logger.Logger.Debug("Getting player news", "user_id", c.GetString("userID"), "ip", c.ClientIP())
	news, err := manager.GetPlayerNews(c.GetString("userID"))
	switch err {
	case nil:
		logger.Logger.Debug("Player news retrieved successfully", "user_id", c.GetString("userID"), "ip", c.ClientIP())
		c.JSON(http.StatusOK, gin.H{
			"news": news,
		})
		return
	case dbErrors.ErrNotFound:
		logger.Logger.Error("Player news not found", "user_id", c.GetString("userID"), "ip", c.ClientIP(), "error", err.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	default:
		logger.Logger.Error("Failed to get player news", "user_id", c.GetString("userID"), "ip", c.ClientIP(), "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

}

func GetCurrentNews(c *gin.Context, manager *postgres.Manager) {
	logger.Logger.Debug("Getting current news", "user_id", c.GetString("userID"), "ip", c.ClientIP())
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Logger.Error("Failed to convert ID", "user_id", c.GetString("userID"), "ip", c.ClientIP(), "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	news, err := manager.GetCurrentNews(id, c.GetString("userID"))
	switch err {
	case nil:
		logger.Logger.Debug("Current news retrieved successfully", "user_id", c.GetString("userID"), "ip", c.ClientIP())
		c.JSON(http.StatusOK, gin.H{
			"news": news,
		})
		return
	case dbErrors.ErrNotFound:
		logger.Logger.Error("Current news not found", "user_id", c.GetString("userID"), "ip", c.ClientIP(), "error", err.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	default:
		logger.Logger.Error("Failed to get current news", "user_id", c.GetString("userID"), "ip", c.ClientIP(), "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
}
