package auth

import "time"

type RefreshToken struct {
	ID        string `json:"id"`
	Token     string `json:"token"`
	ExpiredAt time.Time `json:"expired_at"`
}

type Tokens struct {
	AccessToken  string `json:"access_token" binding:"required"`
	RefreshToken string `json:"refresh_token" binding:"omitempty"`
}