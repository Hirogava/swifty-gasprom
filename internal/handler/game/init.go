package game

import (
	"net/http"

	dbErrors "github.com/Hirogava/swifty-gasprom/backend/internal/errors/db"
	"github.com/Hirogava/swifty-gasprom/backend/internal/handler/middleware"
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
		v1.POST("/life-market/:id", func(c *gin.Context) {
			BuyLifeMarketItem(c, manager)
		})
		v1.GET("/life-market/:id", func(c *gin.Context) {
			GetLifeMarketItem(c, manager)
		})
		v1.GET("/life-market", func(c *gin.Context) {
			GetLifeMarketItems(c, manager)
		})
		v1.GET("/event", func(c *gin.Context) {
			GetRandomEvent(c, manager)
		})
		v1.GET("/vacancies", func(c *gin.Context) {
			GetVacancies(c, manager)
		})
		v1.GET("/vacancy", func(c *gin.Context) {
			GetVacancy(c, manager)
		})
		v1.POST("/vacancy/:id", func(c *gin.Context) {
			SetPlayerVacancy(c, manager)
		})
		v1.PUT("/vacancy/:id", func(c *gin.Context) {
			UpdatePlayerVacancy(c, manager)
		})
		v1.DELETE("/vacancy/:id", func(c *gin.Context) {
			DeletePlayerVacancy(c, manager)
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
