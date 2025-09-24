package bank

import (
	"github.com/Hirogava/swifty-gasprom/backend/internal/handler/middleware"
	"github.com/Hirogava/swifty-gasprom/backend/internal/repository/postgres"

	"github.com/gin-gonic/gin"
)

func InitBankHandlers(r *gin.Engine, manager *postgres.Manager) {
	r.Use(middleware.AuthMiddleware())
	v1 := r.Group("/api/v1")
	{
		v1.POST("/card/:id", func(c *gin.Context) {
			SaveUserBankProduct(c, manager)
		})
		v1.GET("/card", func(c *gin.Context) {
			GetUserBankProduct(c, manager)
		})
		v1.POST("/bonus/:id", func(c *gin.Context) {
			SaveUserBankBonus(c, manager)
		})
		v1.GET("/bonus", func(c *gin.Context) {
			GetUserBankBonuses(c, manager)
		})
		v1.PUT("/bonus/:id", func(c *gin.Context) {
			EditUserBankBonus(c, manager)
		})
	}
}
