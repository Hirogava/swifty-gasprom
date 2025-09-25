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

func GetVacancies(c *gin.Context, manager *postgres.Manager) {
	vacancies, err := manager.GetVacancies()
	switch err {
	case nil:
		c.JSON(http.StatusOK, gin.H{
			"vacancies": vacancies,
		})
		return
	default:
		c.JSON(http.StatusInternalServerError, gin.H{
			"": err.Error(),
		})
	}
}

func GetVacancy(c *gin.Context, manager *postgres.Manager) {
	var req gameModels.GetVacancyRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	vacancy, err := manager.GetVacancy(req.Name)
	switch err {
	case nil:
		c.JSON(http.StatusOK, gin.H{
			"vacancy": vacancy,
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

func GetPlayerVacancy(c *gin.Context, manager *postgres.Manager) {
	userID := c.GetString("userID")

	vacancy, err := manager.GetPlayerVacancy(userID)
	switch err {
	case nil:
		c.JSON(http.StatusOK, gin.H{
			"vacancy": vacancy,
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

func SetPlayerVacancy(c *gin.Context, manager *postgres.Manager) {
	userID := c.GetString("userID")

	var req gameModels.SetVacancyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := manager.SetPlayerVacancy(userID, &req)
	switch err {
	case nil:
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
		})
		return
	default:
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
}

func UpdatePlayerVacancy(c *gin.Context, manager *postgres.Manager) {
	name := c.Param("name")

	var oldVacancy gameModels.UpdatePlayerVacancyRequest

	if err := c.ShouldBindJSON(&oldVacancy); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID := c.GetString("userID")
	vacancy, err := manager.UpdatePlayerVacancy(userID, name, &oldVacancy)
	switch err {
	case nil:
		c.JSON(http.StatusOK, gin.H{
			"vacancy": vacancy,
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

func DeletePlayerVacancy(c *gin.Context, manager *postgres.Manager) {
	vacancyID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = manager.DeletePlayerVacancy(c.GetString("userID"), vacancyID)
	switch err {
	case nil:
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
		})
		return
	case gameErrors.ErrPlayerVacancyNotFound:
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

func UpdatePlayerVacancyGrade(c *gin.Context, manager *postgres.Manager) {
	oldId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	newSalary, penaltyFactor, err := manager.UpdatePlayerVacancyGrade(c.GetString("userID"), oldId)
	switch err {
	case nil:
		c.JSON(http.StatusOK, gin.H{
			"newSalary": newSalary,
			"penaltyFactor": penaltyFactor,
		})
		return
	default:
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
}
