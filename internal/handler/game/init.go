package game

import (
	"net/http"

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
