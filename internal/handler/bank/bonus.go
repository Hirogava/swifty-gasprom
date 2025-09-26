package bank

import (
	"net/http"
	"strconv"

	"github.com/Hirogava/swifty-gasprom/backend/internal/config/logger"
	dbErrors "github.com/Hirogava/swifty-gasprom/backend/internal/errors/db"
	"github.com/Hirogava/swifty-gasprom/backend/internal/repository/postgres"
	bankModels "github.com/Hirogava/swifty-gasprom/backend/internal/models/bank"

	"github.com/gin-gonic/gin"
)

func SaveUserBankBonus(c *gin.Context, manager *postgres.Manager) {
	logger.Logger.Debug("Saving user bank bonus", "user_id", c.GetString("userID"), "ip", c.ClientIP())
	bonusId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Logger.Error("Failed to convert ID", "user_id", c.GetString("userID"), "ip", c.ClientIP(), "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userId := c.GetString("userID")
	logger.Logger.Debug("Saving user bank bonus", "user_id", userId, "bonus_id", bonusId, "ip", c.ClientIP())
	bonus, err := manager.SaveUserBankBonus(userId, bonusId)
	switch err {
	case nil:
		logger.Logger.Debug("User bank bonus saved successfully", "user_id", userId, "bonus_id", bonusId, "ip", c.ClientIP())
		c.JSON(http.StatusOK, gin.H{
			"bonus": bonus,
		})
		return
	case dbErrors.ErrNotFound:
		logger.Logger.Error("User bank bonus not found", "user_id", userId, "bonus_id", bonusId, "ip", c.ClientIP(), "error", err.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	default:
		logger.Logger.Error("Failed to save user bank bonus", "user_id", userId, "bonus_id", bonusId, "ip", c.ClientIP(), "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
}

func GetUserBankBonuses(c *gin.Context, manager *postgres.Manager) {
	bonuses, err := manager.GetUserBankBonuses(c.GetString("userID"))
	switch err {
	case nil:
		logger.Logger.Debug("User bank bonuses retrieved successfully", "user_id", c.GetString("userID"), "ip", c.ClientIP())
		c.JSON(http.StatusOK, gin.H{
			"bonuses": bonuses,
		})
		return
	case dbErrors.ErrNotFound:
		logger.Logger.Error("User bank bonuses not found", "user_id", c.GetString("userID"), "ip", c.ClientIP(), "error", err.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
	default:
		logger.Logger.Error("Failed to get user bank bonuses", "user_id", c.GetString("userID"), "ip", c.ClientIP(), "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
}

func EditUserBankBonus(c *gin.Context, manager *postgres.Manager) {
	var req bankModels.UpdateBonusStatusRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Logger.Error("Failed to bind JSON", "user_id", c.GetString("userID"), "ip", c.ClientIP(), "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := manager.UpdateUserBonusStatusType(req); err != nil {
		logger.Logger.Error("Failed to update user bonus status type", "user_id", c.GetString("userID"), "ip", c.ClientIP(), "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	logger.Logger.Debug("User bonus status type updated successfully", "user_id", c.GetString("userID"), "ip", c.ClientIP())
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}
