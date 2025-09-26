package game

import (
	"net/http"
	"strconv"

	dbErrors "github.com/Hirogava/swifty-gasprom/backend/internal/errors/db"
	gameErrors "github.com/Hirogava/swifty-gasprom/backend/internal/errors/game"
	gameModels "github.com/Hirogava/swifty-gasprom/backend/internal/models/game"
	"github.com/Hirogava/swifty-gasprom/backend/internal/repository/postgres"
	gameService "github.com/Hirogava/swifty-gasprom/backend/internal/service/game"

	"github.com/gin-gonic/gin"
)

func BuyRiskItem(c *gin.Context, manager *postgres.Manager) {
	var envelope gameModels.Envelope

	if err := c.ShouldBindJSON(&envelope); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	msg, err := gameService.EnvelopeRiskStruct(&envelope)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	switch m := msg.(type) {
	case *gameModels.BetsRequest:
		err = manager.BuyBetsItem(m, c.GetString("userID"))
	case *gameModels.CryptoRequest:
		err = manager.BuyCryptoItem(m, c.GetString("userID"))
	case *gameModels.StocksRequest:
		err = manager.BuyStocksItem(m, c.GetString("userID"))
	case *gameModels.QuestionableProjectsRequest:
		err = manager.BuyQPItem(m, c.GetString("userID"))
	}
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

func GetRiskItems(c *gin.Context, manager *postgres.Manager) {
	risk, err := manager.GetRiskItems(c.GetString("userID"))
	switch err {
	case nil:
		c.JSON(http.StatusOK, gin.H{
			"risk": risk,
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

func GetPlayerRiskItems(c *gin.Context, manager *postgres.Manager) {
	items, err := manager.GetPlayerRiskItems(c.GetString("userID"))
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
	}
}

func GetRiskItem(c *gin.Context, manager *postgres.Manager) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	item, err := manager.GetRiskItem(id, c.GetString("userID"), c.Param("type"))
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
	}
}

func GetRiskItemsByType(c *gin.Context, manager *postgres.Manager) {
	items, err := manager.GetRiskItemsByType(c.Param("type"), c.GetString("userID"))
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
	}
}
