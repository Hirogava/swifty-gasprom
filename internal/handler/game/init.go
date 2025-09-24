package game

import (
	"net/http"
	"strconv"

	dbErrors "github.com/Hirogava/swifty-gasprom/backend/internal/errors/db"
	gameErrors "github.com/Hirogava/swifty-gasprom/backend/internal/errors/game"
	"github.com/Hirogava/swifty-gasprom/backend/internal/handler/middleware"
	gameModels "github.com/Hirogava/swifty-gasprom/backend/internal/models/game"
	"github.com/Hirogava/swifty-gasprom/backend/internal/repository/postgres"
	service "github.com/Hirogava/swifty-gasprom/backend/internal/service/game"

	"github.com/gin-gonic/gin"
)

func InitGameHandlers(r *gin.Engine, manager *postgres.Manager) {
	r.Use(middleware.AuthMiddleware())
	v1 := r.Group("/api/v1")
	{
		v1.GET("/start", func(c *gin.Context) {
			StartGame(c, manager)
		})
		v1.GET("/progress", func(c *gin.Context) {
			GetUserProgress(c, manager)
		})
		v1.POST("/life-market/buy/:id", func(c *gin.Context) {
			BuyLifeMarketItem(c, manager)
		})
		v1.GET("/life-market/buy/:id", func(c *gin.Context) {
			GetLifeMarketItem(c, manager)
		})
		v1.GET("/life-market", func(c *gin.Context) {
			GetLifeMarketItems(c, manager)
		})
		v1.GET("/event", func(c *gin.Context) {
			GetRandomEvent(c, manager)
		})
	}
}

func StartGame(c *gin.Context, manager *postgres.Manager) {
	userID := c.Param("userID")

	game := service.StartGame(userID)

	if err := manager.SaveGame(game); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"game": game,
	})
}

func GetUserProgress(c *gin.Context, manager *postgres.Manager) {
	userID := c.Param("userID")

	game, err := manager.GetUserGameInfo(userID)
	switch err {
	case nil:
		c.JSON(http.StatusOK, gin.H{
			"game": game,
		})
	case dbErrors.ErrNotFound:
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Game not found",
		})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
}

func BuyLifeMarketItem(c *gin.Context, manager *postgres.Manager) {
	var item gameModels.LifeMarketItemRequest

	err := c.ShouldBindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	userID := c.Param("userID")

	err = manager.BuyLifeMarketItem(&item, userID)
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
