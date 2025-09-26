package game

import (
	"net/http"
	"strconv"

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
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = manager.BuyLifeMarketItem(&item, c.GetString("userID"))
	switch err {
	case nil:
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
		})
		return
	case gameErrors.ErrNotEnoughMoney:
		c.JSON(http.StatusBadRequest, gin.H{
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

func GetLifeMarketItems(c *gin.Context, manager *postgres.Manager) {
	items, err := manager.GetLifeMarketItems()
	switch err {
	case nil:
		c.JSON(http.StatusOK, gin.H{
			"items": items,
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

func GetLifeMarketItem(c *gin.Context, manager *postgres.Manager) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	item, err := manager.GetLifeMarketItem(id)
	switch err {
	case nil:
		c.JSON(http.StatusOK, gin.H{
			"item": item,
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
