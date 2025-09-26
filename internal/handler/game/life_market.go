package game

import (
	"net/http"
	"strconv"

	"github.com/Hirogava/swifty-gasprom/backend/internal/config/logger"
	dbErrors "github.com/Hirogava/swifty-gasprom/backend/internal/errors/db"
	gameErrors "github.com/Hirogava/swifty-gasprom/backend/internal/errors/game"
	gameModels "github.com/Hirogava/swifty-gasprom/backend/internal/models/game"
	"github.com/Hirogava/swifty-gasprom/backend/internal/repository/postgres"
	
	"github.com/gin-gonic/gin"
)

func BuyLifeMarketItem(c *gin.Context, manager *postgres.Manager) {
	var item gameModels.LifeMarketItemRequest

	err := c.ShouldBindJSON(&item)
	if err != nil {
		logger.Logger.Error("Failed to bind JSON", "user_id", c.GetString("userID"), "ip", c.ClientIP(), "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = manager.BuyLifeMarketItem(&item, c.GetString("userID"))
	switch err {
	case nil:
		logger.Logger.Debug("Life market item bought successfully", "user_id", c.GetString("userID"), "ip", c.ClientIP())
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
		})
		return
	case gameErrors.ErrNotEnoughMoney:
		logger.Logger.Error("Not enough money", "user_id", c.GetString("userID"), "ip", c.ClientIP(), "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	default:
		logger.Logger.Error("Failed to buy life market item", "user_id", c.GetString("userID"), "ip", c.ClientIP(), "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
}

func GetLifeMarketItems(c *gin.Context, manager *postgres.Manager) {
	logger.Logger.Debug("Getting life market items", "user_id", c.GetString("userID"), "ip", c.ClientIP())
	items, err := manager.GetLifeMarketItems()
	switch err {
	case nil:
		logger.Logger.Debug("Life market items retrieved successfully", "user_id", c.GetString("userID"), "ip", c.ClientIP())
		c.JSON(http.StatusOK, gin.H{
			"items": items,
		})
		return
	case dbErrors.ErrNotFound:
		logger.Logger.Error("Life market items not found", "user_id", c.GetString("userID"), "ip", c.ClientIP(), "error", err.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	default:
		logger.Logger.Error("Failed to get life market items", "user_id", c.GetString("userID"), "ip", c.ClientIP(), "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
}

func GetLifeMarketItem(c *gin.Context, manager *postgres.Manager) {
	logger.Logger.Debug("Getting life market item", "user_id", c.GetString("userID"), "ip", c.ClientIP())
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Logger.Error("Failed to convert ID", "user_id", c.GetString("userID"), "ip", c.ClientIP(), "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	item, err := manager.GetLifeMarketItem(id)
	switch err {
	case nil:
		logger.Logger.Debug("Life market item retrieved successfully", "user_id", c.GetString("userID"), "ip", c.ClientIP())
		c.JSON(http.StatusOK, gin.H{
			"item": item,
		})
		return
	case dbErrors.ErrNotFound:
		logger.Logger.Error("Life market item not found", "user_id", c.GetString("userID"), "ip", c.ClientIP(), "error", err.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	default:
		logger.Logger.Error("Failed to get life market item", "user_id", c.GetString("userID"), "ip", c.ClientIP(), "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
}
