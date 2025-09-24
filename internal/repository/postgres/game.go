package postgres

import (
	"database/sql"

	gameModels "github.com/Hirogava/swifty-gasprom/backend/internal/models/game"
	dbErrors "github.com/Hirogava/swifty-gasprom/backend/internal/errors/db"
	gameErrors "github.com/Hirogava/swifty-gasprom/backend/internal/errors/game"
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

func (manager *Manager) GetRandomEvent() (*gameModels.Event, error) {
	var event gameModels.Event

	if err := manager.Conn.QueryRow(`SELECT id, name, event_text, capital_percent, to_agree, refuse FROM events ORDER BY RANDOM() LIMIT 1`).Scan(&event.ID, &event.Name, &event.EventText, &event.CapitalPercent, &event.ToAgree, &event.Refuse); err != nil {
		return nil, err
	} 

	return &event, nil
}

func (manager *Manager) GetLifeMarketItem(id int) (*gameModels.LifeMarketItem, error) {
	var item gameModels.LifeMarketItem

	if err := manager.Conn.QueryRow(`SELECT id, name, cost, category, effect_duration FROM life_market WHERE id = $1`, id).Scan(&item.ID, &item.Name, &item.Cost, &item.Category, &item.HappinessEffect); err != nil {
		if err == sql.ErrNoRows {
			return nil, dbErrors.ErrNotFound
		} else {
			return nil, err
		}
	}

	return &item, nil
}

func (manager *Manager) GetLifeMarketItems() ([]gameModels.LifeMarketItem, error) {
	var items []gameModels.LifeMarketItem

	rows, err := manager.Conn.Query(`SELECT id, name, cost, category, effect_duration FROM life_market`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var item gameModels.LifeMarketItem

		if err := rows.Scan(&item.ID, &item.Name, &item.Cost, &item.Category, &item.HappinessEffect); err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
}

func (manager *Manager) BuyLifeMarketItem(req *gameModels.LifeMarketItemRequest, userID string) error {
	tx, err := manager.Conn.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(`INSERT INTO user_assets (user_id, market_id) VALUES ($1, $2)`, userID, req.ID)
	if err != nil {
		return err
	}

	var money float64
	err = tx.QueryRow(`UPDATE user_progress SET money = money - $1, happiness = happiness + $2 WHERE user_id = $3 RETURNING money`, req.Cost, req.Happiness, userID).Scan(&money)
	if err != nil {
		return err
	}

	if money < 0 {
		return gameErrors.ErrNotEnoughMoney
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	
	return nil
}
