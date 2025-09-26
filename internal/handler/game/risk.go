package game

import (
	"net/http"
	"strconv"

	"github.com/Hirogava/swifty-gasprom/backend/internal/config/logger"
	dbErrors "github.com/Hirogava/swifty-gasprom/backend/internal/errors/db"
	gameErrors "github.com/Hirogava/swifty-gasprom/backend/internal/errors/game"
	gameModels "github.com/Hirogava/swifty-gasprom/backend/internal/models/game"
	"github.com/Hirogava/swifty-gasprom/backend/internal/repository/postgres"
	gameService "github.com/Hirogava/swifty-gasprom/backend/internal/service/game"

	"github.com/gin-gonic/gin"
)

func BuyRiskItem(c *gin.Context, manager *postgres.Manager) {
	logger.Logger.Debug("Buying risk item", "user_id", c.GetString("userID"), "ip", c.ClientIP())
	var envelope gameModels.Envelope

	if err := c.ShouldBindJSON(&envelope); err != nil {
		logger.Logger.Error("Failed to bind JSON", "user_id", c.GetString("userID"), "ip", c.ClientIP(), "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	msg, err := gameService.EnvelopeRiskStruct(&envelope)
	if err != nil {
		logger.Logger.Error("Failed to envelope risk struct", "user_id", c.GetString("userID"), "ip", c.ClientIP(), "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	switch m := msg.(type) {
	case *gameModels.BetsRequest:
		err = manager.BuyBetsItem(m, c.GetString("userID"))
	case *gameModels.CryptoRequest:
		err = manager.BuyCryptoItem(m, c.GetString("userID"))
	case *gameModels.StocksRequest:
		err = manager.BuyStocksItem(m, c.GetString("userID"))
	case *gameModels.QuestionableProjectsRequest:
		err = manager.BuyQPItem(m, c.GetString("userID"))
	}
	switch err {
	case nil:
		logger.Logger.Debug("Risk item bought successfully", "user_id", c.GetString("userID"), "ip", c.ClientIP())
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
	case dbErrors.ErrNotFound:
		logger.Logger.Error("Risk item not found", "user_id", c.GetString("userID"), "ip", c.ClientIP(), "error", err.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	default:
		logger.Logger.Error("Failed to buy risk item", "user_id", c.GetString("userID"), "ip", c.ClientIP(), "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
}

func GetRiskItems(c *gin.Context, manager *postgres.Manager) {
	logger.Logger.Debug("Getting risk items", "user_id", c.GetString("userID"), "ip", c.ClientIP())
	risk, err := manager.GetRiskItems(c.GetString("userID"))
	switch err {
	case nil:
		logger.Logger.Debug("Risk items retrieved successfully", "user_id", c.GetString("userID"), "ip", c.ClientIP())
		c.JSON(http.StatusOK, gin.H{
			"risk": risk,
		})
		return
	case dbErrors.ErrNotFound:
		logger.Logger.Error("Risk items not found", "user_id", c.GetString("userID"), "ip", c.ClientIP(), "error", err.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	default:
		logger.Logger.Error("Failed to get risk items", "user_id", c.GetString("userID"), "ip", c.ClientIP(), "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
}

func GetPlayerRiskItems(c *gin.Context, manager *postgres.Manager) {
	logger.Logger.Debug("Getting player risk items", "user_id", c.GetString("userID"), "ip", c.ClientIP())
	items, err := manager.GetPlayerRiskItems(c.GetString("userID"))
	switch err {
	case nil:
		logger.Logger.Debug("Player risk items retrieved successfully", "user_id", c.GetString("userID"), "ip", c.ClientIP())
		c.JSON(http.StatusOK, gin.H{
			"items": items,
		})
		return
	case dbErrors.ErrNotFound:
		logger.Logger.Error("Player risk items not found", "user_id", c.GetString("userID"), "ip", c.ClientIP(), "error", err.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	default:
		logger.Logger.Error("Failed to get player risk items", "user_id", c.GetString("userID"), "ip", c.ClientIP(), "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
}

func GetRiskItem(c *gin.Context, manager *postgres.Manager) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Logger.Error("Failed to convert ID", "user_id", c.GetString("userID"), "ip", c.ClientIP(), "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	item, err := manager.GetRiskItem(id, c.GetString("userID"), c.Param("type"))
	switch err {
	case nil:
		logger.Logger.Debug("Risk item retrieved successfully", "user_id", c.GetString("userID"), "ip", c.ClientIP())
		c.JSON(http.StatusOK, gin.H{
			"item": item,
		})
		return
	case dbErrors.ErrNotFound:
		logger.Logger.Error("Risk item not found", "user_id", c.GetString("userID"), "ip", c.ClientIP(), "error", err.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	default:
		logger.Logger.Error("Failed to get risk item", "user_id", c.GetString("userID"), "ip", c.ClientIP(), "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
}

func GetRiskItemsByType(c *gin.Context, manager *postgres.Manager) {
	logger.Logger.Debug("Getting risk items by type", "user_id", c.GetString("userID"), "ip", c.ClientIP())
	items, err := manager.GetRiskItemsByType(c.Param("type"), c.GetString("userID"))
	switch err {
	case nil:
		logger.Logger.Debug("Risk items by type retrieved successfully", "user_id", c.GetString("userID"), "ip", c.ClientIP())
		c.JSON(http.StatusOK, gin.H{
			"items": items,
		})
		return
	case dbErrors.ErrNotFound:
		logger.Logger.Error("Risk items by type not found", "user_id", c.GetString("userID"), "ip", c.ClientIP(), "error", err.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	default:
		logger.Logger.Error("Failed to get risk items by type", "user_id", c.GetString("userID"), "ip", c.ClientIP(), "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
}
