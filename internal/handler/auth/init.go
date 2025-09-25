package auth

import (
	"net/http"

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
	var req authModels.BankUser 

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Псевдо-логика проверки данных
	user, err := manager.FindOrCreateUser(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Set("userID", user.ID)

	var refreshToken string
	token, err := tokens.ValidateRefreshToken(manager, user.ID)
	if err != nil {
		if err == jwt.ErrTokenExpired {
			refreshToken, err = tokens.GenerateRefreshToken(manager, user.ID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		} else {
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
		c.JSON(http.StatusInternalServerError, gin.H{"": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user" : user,
	})
}

func RefreshToken(c *gin.Context, manager *postgres.Manager) {
	var t authModels.Tokens
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId := c.GetString("userID")

	refreshToken, err := tokens.ValidateRefreshToken(manager, userId)
	if err != nil {
		switch err {
			case jwt.ErrTokenExpired:
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token expired"})
				return
			default:
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": t.AccessToken,
		"refresh_token": t.RefreshToken,
	})
}

func Logout(c *gin.Context, manager *postgres.Manager) {
	userId := c.GetString("userID")

	token, err := tokens.ValidateRefreshToken(manager, userId)
	if err != nil {
		switch err {
			case jwt.ErrTokenExpired:
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token expired"})
				return
			default:
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
		}
	}

	if err := manager.DeleteRefreshToken(userId, token.Token); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"logout": "success",
	})
}
