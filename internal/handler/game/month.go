package game

import (
	"net/http"

	"github.com/Hirogava/swifty-gasprom/backend/internal/repository/postgres"

	"github.com/gin-gonic/gin"
)

func NextMonthMove(c *gin.Context, manager *postgres.Manager) {
	newMonth, err := manager.NextMonthMove(c.GetString("userID"))
	switch err {
	case nil:
		c.JSON(http.StatusOK, gin.H{
			"newMonth": newMonth,
		})
		return
	default:
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
}
