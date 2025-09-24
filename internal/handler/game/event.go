package game

import (
	"net/http"

	"github.com/Hirogava/swifty-gasprom/backend/internal/repository/postgres"
	
	"github.com/gin-gonic/gin"
)

func GetRandomEvent(c *gin.Context, manager *postgres.Manager) {
	event, err := manager.GetRandomEvent()
	switch err {
	case nil:
		c.JSON(http.StatusOK, gin.H{
			"event": event,
		})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
}