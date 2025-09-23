package postgres

import (
	gameModels "github.com/Hirogava/swifty-gasprom/backend/internal/models/game"
)

func (manager *Manager) SaveGame(game *gameModels.UserProgress) error {
	err := manager.Conn.QueryRow(`INSERT INTO user_progress (user_id, money, happiness, month, status) VALUES ($1, $2, $3, $4, $5)`, game.UserID, game.Money, game.Happiness, game.Month, game.Status).Scan(&game.ID)
	if err != nil {
		return err
	}

	return nil
}