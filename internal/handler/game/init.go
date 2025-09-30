package game

import (
	"net/http"

	"github.com/Hirogava/swifty-gasprom/backend/internal/config/logger"
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
		v1.GET("/player/vacancy", func(c *gin.Context) {
			GetPlayerVacancy(c, manager)
		})
		v1.POST("/player/vacancy", func(c *gin.Context) {
			SetPlayerVacancy(c, manager)
		})
		v1.PUT("/player/vacancy/:name", func(c *gin.Context) {
			UpdatePlayerVacancy(c, manager)
		})
		v1.DELETE("/player/vacancy/:id", func(c *gin.Context) {
			DeletePlayerVacancy(c, manager)
		})
		v1.PUT("/player/vacancy/grade/:id", func(c *gin.Context) {
			UpdatePlayerVacancyGrade(c, manager)
		})
		v1.POST("/risk", func(c *gin.Context) {
			BuyRiskItem(c, manager)
		})
		v1.GET("/risk", func(c *gin.Context) {
			GetRiskItems(c, manager)
		})
		v1.GET("/risk/type/:type", func(c *gin.Context) {
			GetRiskItemsByType(c, manager)
		})
		v1.GET("/risk/:id/:riskType", func(c *gin.Context) {
			GetRiskItem(c, manager)
		})
		v1.GET("/player/risk", func(c *gin.Context) {
			GetPlayerRiskItems(c, manager)
		})
		v1.GET("/player/news", func(c *gin.Context) {
			GetPlayerNews(c, manager)
		})
		v1.GET("/player/news/:id", func(c *gin.Context) {
			GetCurrentNews(c, manager)
		})
		v1.POST("/month/next", func(c *gin.Context) {
			NextMonthMove(c, manager)
		})
	}
}

func StartGame(c *gin.Context, manager *postgres.Manager) {
	userID := c.GetString("userID")
	logger.Logger.Info("Starting new game", "user_id", userID, "ip", c.ClientIP())

	game := service.StartGame(userID)
	logger.Logger.Debug("Game object created", "user_id", userID, "money", game.Money, "happiness", game.Happiness)

	if err := manager.SaveGame(game); err != nil {
		logger.Logger.Error("Failed to save game", "user_id", userID, "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	logger.Logger.Debug("Game saved to database", "user_id", userID)

	if err := manager.InitStocksAndCryptos(userID); err != nil {
		logger.Logger.Error("Failed to initialize stocks and cryptos", "user_id", userID, "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	logger.Logger.Debug("Stocks and cryptos initialized", "user_id", userID)

	logger.Logger.Info("Game started successfully", "user_id", userID, "ip", c.ClientIP())

	c.JSON(http.StatusOK, gin.H{
		"game": game,
	})
}

func GetUserProgress(c *gin.Context, manager *postgres.Manager) {
	userID := c.GetString("userID")
	logger.Logger.Debug("Getting user progress", "user_id", userID, "ip", c.ClientIP())

	game, err := manager.GetUserGameInfo(userID)
	switch err {
	case nil:
		logger.Logger.Debug("User progress retrieved successfully", "user_id", userID, "money", game.UserProgress.Money, "happiness", game.UserProgress.Happiness, "month", game.UserProgress.Month)
		c.JSON(http.StatusOK, gin.H{
			"game": game,
		})
		return
	case dbErrors.ErrNotFound:
		logger.Logger.Warn("Game not found for user", "user_id", userID, "ip", c.ClientIP())
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Game not found",
		})
		return
	default:
		logger.Logger.Error("Failed to get user game info", "user_id", userID, "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
}
