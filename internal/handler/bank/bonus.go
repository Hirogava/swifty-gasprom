package bank

import (
	"net/http"
	"strconv"

	dbErrors "github.com/Hirogava/swifty-gasprom/backend/internal/errors/db"
	"github.com/Hirogava/swifty-gasprom/backend/internal/repository/postgres"
	bankModels "github.com/Hirogava/swifty-gasprom/backend/internal/models/bank"

	"github.com/gin-gonic/gin"
)

func SaveUserBankBonus(c *gin.Context, manager *postgres.Manager) {
	bonusId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userId := c.GetString("userID")

	bonus, err := manager.SaveUserBankBonus(userId, bonusId)
	switch err {
	case nil:
		c.JSON(http.StatusOK, gin.H{
			"bonus": bonus,
		})
		return
	case dbErrors.ErrNotFound:
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	default:
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
		c.JSON(http.StatusOK, gin.H{
			"bonuses": bonuses,
		})
		return
	case dbErrors.ErrNotFound:
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
}

func EditUserBankBonus(c *gin.Context, manager *postgres.Manager) {
	var req bankModels.UpdateBonusStatusRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := manager.UpdateUserBonusStatusType(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}
