package postgres

import (
	"database/sql"

	gameModels "github.com/Hirogava/swifty-gasprom/backend/internal/models/game"
	dbErrors "github.com/Hirogava/swifty-gasprom/backend/internal/errors/db"
)

func (manager *Manager) SaveGame(game *gameModels.UserProgress) error {
	err := manager.Conn.QueryRow(`INSERT INTO user_progress (user_id, money, happiness, month, status) VALUES ($1, $2, $3, $4, $5)`, game.UserID, game.Money, game.Happiness, game.Month, game.Status).Scan(&game.ID)
	if err != nil {
		return err
	}

	return nil
}

func (manager *Manager) GetUserGameInfo(userID string) (*gameModels.UserGameInfo, error) {
	var gameInfo gameModels.UserGameInfo

	err := manager.Conn.QueryRow(`SELECT 
			up.id AS user_progress_id,
			up.user_id,
			up.money,
			up.happiness,
			up.month,
			up.status,

			n.id AS news_id,
			n.news_title,
			n.news_text,

			c.id AS career_id,
			uc.career_level,
			c.name AS career_name,
			cf.name AS career_field
		FROM user_progress up
		LEFT JOIN current_progress_news cpn 
			ON up.user_id = cpn.user_id
		LEFT JOIN news n 
			ON cpn.news_id = n.id
		LEFT JOIN user_career uc 
			ON up.user_id = uc.user_id
		LEFT JOIN careers c 
			ON uc.career_id = c.id
		LEFT JOIN career_fields cf 
			ON c.field_id = cf.id
		WHERE up.user_id = $1
		`).Scan(&gameInfo.UserProgress.ID, &gameInfo.UserProgress.UserID, &gameInfo.UserProgress.Money,
		&gameInfo.UserProgress.Happiness, &gameInfo.UserProgress.Month, &gameInfo.UserProgress.Status,
		&gameInfo.News.ID, &gameInfo.News.Title, &gameInfo.News.Text, &gameInfo.Career.ID,
		&gameInfo.Career.Level, &gameInfo.Career.Name, &gameInfo.Career.CareerField)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, dbErrors.ErrNotFound
		} else {
			return nil, err
		}
	}

	return &gameInfo, nil
}
