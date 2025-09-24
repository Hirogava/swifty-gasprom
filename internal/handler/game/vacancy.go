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

func GetVacancies(c *gin.Context, manager *postgres.Manager) {}

func GetVacancy(c *gin.Context, manager *postgres.Manager) {}

func SetPlayerVacancy(c *gin.Context, manager *postgres.Manager) {}

func UpdatePlayerVacancy(c *gin.Context, manager *postgres.Manager) {}

func DeletePlayerVacancy(c *gin.Context, manager *postgres.Manager) {}