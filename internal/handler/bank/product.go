package bank

import (
	"net/http"
	"strconv"

	dbErrors "github.com/Hirogava/swifty-gasprom/backend/internal/errors/db"
	"github.com/Hirogava/swifty-gasprom/backend/internal/repository/postgres"
	
	"github.com/gin-gonic/gin"
)

func SaveUserBankProduct(c *gin.Context, manager *postgres.Manager) {
	cardId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userId := c.GetString("userID")

	card, err := manager.SaveUserBankProduct(userId, cardId)
	switch err {
	case nil:
		c.JSON(http.StatusOK, gin.H{
			"card": card,
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

func GetUserBankProduct(c *gin.Context, manager *postgres.Manager) {
	card, err := manager.GetUserBankProduct(c.GetString("userID"))
	switch err {
	case nil:
		c.JSON(http.StatusOK, gin.H{
			"card": card,
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