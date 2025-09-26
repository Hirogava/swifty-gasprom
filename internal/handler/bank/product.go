package bank

import (
	"net/http"
	"strconv"

	"github.com/Hirogava/swifty-gasprom/backend/internal/config/logger"
	dbErrors "github.com/Hirogava/swifty-gasprom/backend/internal/errors/db"
	"github.com/Hirogava/swifty-gasprom/backend/internal/repository/postgres"

	"github.com/gin-gonic/gin"
)

func SaveUserBankProduct(c *gin.Context, manager *postgres.Manager) {
	userId := c.GetString("userID")

	cardId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Logger.Warn("Invalid card ID parameter", "user_id", userId, "card_id_param", c.Param("id"), "ip", c.ClientIP(), "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	logger.Logger.Info("Saving user bank product", "user_id", userId, "card_id", cardId, "ip", c.ClientIP())

	card, err := manager.SaveUserBankProduct(userId, cardId)
	switch err {
	case nil:
		logger.Logger.Info("Bank product saved successfully", "user_id", userId, "card_id", cardId, "card_name", card.Name)
		c.JSON(http.StatusOK, gin.H{
			"card": card,
		})
		return
	case dbErrors.ErrNotFound:
		logger.Logger.Warn("Bank product not found", "user_id", userId, "card_id", cardId, "ip", c.ClientIP())
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	default:
		logger.Logger.Error("Failed to save bank product", "user_id", userId, "card_id", cardId, "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
}

func GetUserBankProduct(c *gin.Context, manager *postgres.Manager) {
	userId := c.GetString("userID")
	logger.Logger.Debug("Getting user bank product", "user_id", userId, "ip", c.ClientIP())

	card, err := manager.GetUserBankProduct(userId)
	switch err {
	case nil:
		logger.Logger.Debug("Bank product retrieved successfully", "user_id", userId, "card_id", card.ID, "card_name", card.Name)
		c.JSON(http.StatusOK, gin.H{
			"card": card,
		})
		return
	case dbErrors.ErrNotFound:
		logger.Logger.Debug("No bank product found for user", "user_id", userId, "ip", c.ClientIP())
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	default:
		logger.Logger.Error("Failed to get user bank product", "user_id", userId, "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
}
