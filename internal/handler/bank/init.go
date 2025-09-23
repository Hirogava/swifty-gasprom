package bank

import (
	"net/http"
	"strconv"

	dbErrors "github.com/Hirogava/swifty-gasprom/backend/internal/errors/db"
	bankModels "github.com/Hirogava/swifty-gasprom/backend/internal/models/bank"
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

func SaveUserBankBonus(c *gin.Context, manager *postgres.Manager) {
	bonusId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userId := c.GetString("userID")

	bonus, err := manager.SaveUserBankBonus(userId, bonusId)
	switch err {
	case nil: 
		c.JSON(http.StatusOK, gin.H{
			"bonus": bonus,
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

func GetUserBankBonuses(c *gin.Context, manager *postgres.Manager) {
	bonuses, err := manager.GetUserBankBonuses(c.GetString("userID"))
	switch err {
	case nil:
		c.JSON(http.StatusOK, gin.H{
			"bonuses": bonuses,
		})
		return
	case dbErrors.ErrNotFound:
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
}

func EditUserBankBonus(c *gin.Context, manager *postgres.Manager) {
	var req bankModels.UpdateBonusStatusRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	if err := manager.UpdateUserBonusStatusType(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}
