package auth

import (
	"github.com/Hirogava/swifty-gasprom/backend/internal/repository/postgres"
	
	"github.com/gin-gonic/gin"
)

func InitAuthHandlers(r *gin.Engine, manager *postgres.Manager) {}