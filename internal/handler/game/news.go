package game

import (
	"net/http"
	"strconv"

	dbErrors "github.com/Hirogava/swifty-gasprom/backend/internal/errors/db"
	"github.com/Hirogava/swifty-gasprom/backend/internal/repository/postgres"

	"github.com/gin-gonic/gin"
)

func GetPlayerNews(c *gin.Context, manager *postgres.Manager) {
	news, err := manager.GetPlayerNews(c.Param("userID"))
	switch err {
	case nil:
		c.JSON(http.StatusOK, gin.H{
			"news": news,
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
	}

}

func GetCurrentNews(c *gin.Context, manager *postgres.Manager) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	news, err := manager.GetCurrentNews(id, c.Param("userID"))
	switch err {
	case nil:
		c.JSON(http.StatusOK, gin.H{
			"news": news,
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
	}
}
