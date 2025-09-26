package auth

import (
	"net/http"

	"github.com/Hirogava/swifty-gasprom/backend/internal/config/logger"
	"github.com/Hirogava/swifty-gasprom/backend/internal/handler/middleware"
	authModels "github.com/Hirogava/swifty-gasprom/backend/internal/models/auth"
	"github.com/Hirogava/swifty-gasprom/backend/internal/repository/postgres"
	tokens "github.com/Hirogava/swifty-gasprom/backend/internal/service/auth"
	"github.com/golang-jwt/jwt/v5"

	"github.com/gin-gonic/gin"
)

func InitAuthHandlers(r *gin.Engine, manager *postgres.Manager) {
	v1 := r.Group("/api/v1")
	{
		v1.POST("/login", func(c *gin.Context) {
			Login(c, manager)
		})
	}

	secureV1 := r.Group("/api/v1")
	secureV1.Use(middleware.AuthMiddleware())
	{
		secureV1.POST("/refresh", func(c *gin.Context) {
			RefreshToken(c, manager)
		})
		secureV1.POST("/logout", func(c *gin.Context) {
			Logout(c, manager)
		})
	}
}

func Login(c *gin.Context, manager *postgres.Manager) {
	logger.Logger.Info("Login attempt", "ip", c.ClientIP())

	var req authModels.BankUser

	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Logger.Warn("Invalid login request", "ip", c.ClientIP(), "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	logger.Logger.Debug("Processing login for bank user", "bank_user_id", req.BankUserID, "ip", c.ClientIP())

	// Псевдо-логика проверки данных
	user, err := manager.FindOrCreateUser(req)
	if err != nil {
		logger.Logger.Error("Failed to find or create user", "bank_user_id", req.BankUserID, "ip", c.ClientIP(), "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.Logger.Info("User found/created successfully", "user_id", user.ID, "bank_user_id", req.BankUserID, "ip", c.ClientIP())

	c.Set("userID", user.ID)

	var refreshToken string
	token, err := tokens.ValidateRefreshToken(manager, user.ID)
	if err != nil {
		if err == jwt.ErrTokenExpired {
			logger.Logger.Debug("Refresh token expired, generating new one", "user_id", user.ID)
			refreshToken, err = tokens.GenerateRefreshToken(manager, user.ID)
			if err != nil {
				logger.Logger.Error("Failed to generate refresh token", "user_id", user.ID, "error", err.Error())
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		} else {
			logger.Logger.Error("Failed to validate refresh token", "user_id", user.ID, "error", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	if token.Token == "" {
		token.Token = refreshToken
	}

	user.Token.RefreshToken = token.Token

	claims := jwt.MapClaims{
		"id":  user.ID,
		"exp": tokens.AddAccessTime(),
	}

	if user.Token.AccessToken, err = tokens.GenerateAccessToken(claims); err != nil {
		logger.Logger.Error("Failed to generate access token", "user_id", user.ID, "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"": err.Error()})
		return
	}

	logger.Logger.Info("Login successful", "user_id", user.ID, "bank_user_id", req.BankUserID, "ip", c.ClientIP())

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func RefreshToken(c *gin.Context, manager *postgres.Manager) {
	userId := c.GetString("userID")
	logger.Logger.Info("Token refresh attempt", "user_id", userId, "ip", c.ClientIP())

	var t authModels.Tokens
	if err := c.ShouldBindJSON(&t); err != nil {
		logger.Logger.Warn("Invalid refresh token request", "user_id", userId, "ip", c.ClientIP(), "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	refreshToken, err := tokens.ValidateRefreshToken(manager, userId)
	if err != nil {
		switch err {
		case jwt.ErrTokenExpired:
			logger.Logger.Warn("Refresh token expired", "user_id", userId, "ip", c.ClientIP())
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token expired"})
			return
		default:
			logger.Logger.Error("Failed to validate refresh token", "user_id", userId, "ip", c.ClientIP(), "error", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	t.RefreshToken = refreshToken.Token

	claims := jwt.MapClaims{
		"id":  userId,
		"exp": tokens.AddAccessTime(),
	}

	if t.AccessToken, err = tokens.GenerateAccessToken(claims); err != nil {
		logger.Logger.Error("Failed to generate new access token", "user_id", userId, "ip", c.ClientIP(), "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.Logger.Info("Token refresh successful", "user_id", userId, "ip", c.ClientIP())

	c.JSON(http.StatusOK, gin.H{
		"access_token":  t.AccessToken,
		"refresh_token": t.RefreshToken,
	})
}

func Logout(c *gin.Context, manager *postgres.Manager) {
	userId := c.GetString("userID")
	logger.Logger.Info("Logout attempt", "user_id", userId, "ip", c.ClientIP())

	token, err := tokens.ValidateRefreshToken(manager, userId)
	if err != nil {
		switch err {
		case jwt.ErrTokenExpired:
			logger.Logger.Warn("Refresh token expired during logout", "user_id", userId, "ip", c.ClientIP())
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token expired"})
			return
		default:
			logger.Logger.Error("Failed to validate refresh token during logout", "user_id", userId, "ip", c.ClientIP(), "error", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	if err := manager.DeleteRefreshToken(userId, token.Token); err != nil {
		logger.Logger.Error("Failed to delete refresh token", "user_id", userId, "ip", c.ClientIP(), "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.Logger.Info("Logout successful", "user_id", userId, "ip", c.ClientIP())

	c.JSON(http.StatusOK, gin.H{
		"logout": "success",
	})
}
