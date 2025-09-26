package game

import (
	"net/http"
	"strconv"

	"github.com/Hirogava/swifty-gasprom/backend/internal/config/logger"
	dbErrors "github.com/Hirogava/swifty-gasprom/backend/internal/errors/db"
	gameErrors "github.com/Hirogava/swifty-gasprom/backend/internal/errors/game"
	gameModels "github.com/Hirogava/swifty-gasprom/backend/internal/models/game"
	"github.com/Hirogava/swifty-gasprom/backend/internal/repository/postgres"

	"github.com/gin-gonic/gin"
)

func GetVacancies(c *gin.Context, manager *postgres.Manager) {
	logger.Logger.Debug("Getting vacancies", "user_id", c.GetString("userID"), "ip", c.ClientIP())
	vacancies, err := manager.GetVacancies()
	switch err {
	case nil:
		logger.Logger.Debug("Vacancies retrieved successfully", "user_id", c.GetString("userID"), "ip", c.ClientIP())
		c.JSON(http.StatusOK, gin.H{
			"vacancies": vacancies,
		})
		return
	default:
		logger.Logger.Error("Failed to get vacancies", "user_id", c.GetString("userID"), "ip", c.ClientIP(), "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"": err.Error(),
		})
	}
}

func GetVacancy(c *gin.Context, manager *postgres.Manager) {
	logger.Logger.Debug("Getting vacancy", "user_id", c.GetString("userID"), "ip", c.ClientIP())
	var req gameModels.GetVacancyRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Logger.Error("Failed to bind JSON", "user_id", c.GetString("userID"), "ip", c.ClientIP(), "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	vacancy, err := manager.GetVacancy(req.Name)
	switch err {
	case nil:
		logger.Logger.Debug("Vacancy retrieved successfully", "user_id", c.GetString("userID"), "ip", c.ClientIP())
		c.JSON(http.StatusOK, gin.H{
			"vacancy": vacancy,
		})
		return
	case dbErrors.ErrNotFound:
		logger.Logger.Error("Vacancy not found", "user_id", c.GetString("userID"), "ip", c.ClientIP(), "error", err.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	default:
		logger.Logger.Error("Failed to get vacancy", "user_id", c.GetString("userID"), "ip", c.ClientIP(), "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
}

func GetPlayerVacancy(c *gin.Context, manager *postgres.Manager) {
	userID := c.GetString("userID")
	logger.Logger.Debug("Getting player vacancy", "user_id", userID, "ip", c.ClientIP())

	vacancy, err := manager.GetPlayerVacancy(userID)
	switch err {
	case nil:
		logger.Logger.Debug("Player vacancy retrieved successfully", "user_id", userID, "ip", c.ClientIP())
		c.JSON(http.StatusOK, gin.H{
			"vacancy": vacancy,
		})
		return
	case dbErrors.ErrNotFound:
		logger.Logger.Error("Player vacancy not found", "user_id", userID, "ip", c.ClientIP(), "error", err.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	default:
		logger.Logger.Error("Failed to get player vacancy", "user_id", userID, "ip", c.ClientIP(), "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
}

func SetPlayerVacancy(c *gin.Context, manager *postgres.Manager) {
	userID := c.GetString("userID")
	logger.Logger.Debug("Setting player vacancy", "user_id", userID, "ip", c.ClientIP())

	var req gameModels.SetVacancyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Logger.Error("Failed to bind JSON", "user_id", userID, "ip", c.ClientIP(), "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := manager.SetPlayerVacancy(userID, &req)
	switch err {
	case nil:
		logger.Logger.Debug("Player vacancy set successfully", "user_id", userID, "ip", c.ClientIP())
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
		})
		return
	default:
		logger.Logger.Error("Failed to set player vacancy", "user_id", userID, "ip", c.ClientIP(), "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
}

func UpdatePlayerVacancy(c *gin.Context, manager *postgres.Manager) {
	name := c.Param("name")
	logger.Logger.Debug("Updating player vacancy", "user_id", c.GetString("userID"), "ip", c.ClientIP())
	var oldVacancy gameModels.UpdatePlayerVacancyRequest

	if err := c.ShouldBindJSON(&oldVacancy); err != nil {
		logger.Logger.Error("Failed to bind JSON", "user_id", c.GetString("userID"), "ip", c.ClientIP(), "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID := c.GetString("userID")
	vacancy, err := manager.UpdatePlayerVacancy(userID, name, &oldVacancy)
	switch err {
	case nil:
		logger.Logger.Debug("Player vacancy updated successfully", "user_id", c.GetString("userID"), "ip", c.ClientIP())
		c.JSON(http.StatusOK, gin.H{
			"vacancy": vacancy,
		})
		return
	case dbErrors.ErrNotFound:
		logger.Logger.Error("Player vacancy not found", "user_id", c.GetString("userID"), "ip", c.ClientIP(), "error", err.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	default:
		logger.Logger.Error("Failed to update player vacancy", "user_id", c.GetString("userID"), "ip", c.ClientIP(), "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
}

func DeletePlayerVacancy(c *gin.Context, manager *postgres.Manager) {
	logger.Logger.Debug("Deleting player vacancy", "user_id", c.GetString("userID"), "ip", c.ClientIP())
	vacancyID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Logger.Error("Failed to convert ID", "user_id", c.GetString("userID"), "ip", c.ClientIP(), "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = manager.DeletePlayerVacancy(c.GetString("userID"), vacancyID)
	switch err {
	case nil:
		logger.Logger.Debug("Player vacancy deleted successfully", "user_id", c.GetString("userID"), "ip", c.ClientIP())
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
		})
		return
	case gameErrors.ErrPlayerVacancyNotFound:
		logger.Logger.Error("Player vacancy not found", "user_id", c.GetString("userID"), "ip", c.ClientIP(), "error", err.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	default:
		logger.Logger.Error("Failed to delete player vacancy", "user_id", c.GetString("userID"), "ip", c.ClientIP(), "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
}

func UpdatePlayerVacancyGrade(c *gin.Context, manager *postgres.Manager) {
	logger.Logger.Debug("Updating player vacancy grade", "user_id", c.GetString("userID"), "ip", c.ClientIP())
	oldId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Logger.Error("Failed to convert ID", "user_id", c.GetString("userID"), "ip", c.ClientIP(), "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	newSalary, penaltyFactor, err := manager.UpdatePlayerVacancyGrade(c.GetString("userID"), oldId)
	switch err {
	case nil:
		logger.Logger.Debug("Player vacancy grade updated successfully", "user_id", c.GetString("userID"), "ip", c.ClientIP())
		c.JSON(http.StatusOK, gin.H{
			"newSalary": newSalary,
			"penaltyFactor": penaltyFactor,
		})
		return
	default:
		logger.Logger.Error("Failed to update player vacancy grade", "user_id", c.GetString("userID"), "ip", c.ClientIP(), "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
}
