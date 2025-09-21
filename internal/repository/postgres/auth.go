package postgres

import (
	"time"

	authModels "github.com/Hirogava/swifty-gasprom/backend/internal/models/auth"
)

func (manager *Manager) SaveRefreshToken(token string, userId string) error {
	expiredAt := time.Now().Add(time.Hour * 24 * 7)

	_, err := manager.Conn.Exec(`INSERT INTO user_tokens (user_id, token, expires_at) VALUES ($1, $2, $3)`, userId, token, expiredAt)
	return err
}

func (manager *Manager) GetRefreshToken(userId string) (authModels.RefreshToken, error) {
	var token authModels.RefreshToken

	if err := manager.Conn.QueryRow(`SELECT id, token, expires_at FROM user_tokens WHERE user_id = $1`, userId).Scan(&token.ID, &token.Token, &token.ExpiredAt); err != nil {
		return authModels.RefreshToken{}, err
	}

	return token, nil
}

func (manager *Manager) DeleteRefreshToken(userId string, token string) error {
	_, err := manager.Conn.Exec(`DELETE FROM user_tokens WHERE user_id = $1`, userId)

	return err
}

func (manager *Manager) UpdateRefreshToken(userId string, token string) error {
	if _, err := manager.Conn.Exec(`UPDATE user_tokens SET expires_at = $1, token = $2 WHERE user_id = $3`, time.Now().Add(time.Hour*24*7).Unix(), token, userId); err != nil {
		return err
	}

	return nil
}
