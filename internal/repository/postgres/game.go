package postgres

import (
	"database/sql"
	"encoding/json"

	dbErrors "github.com/Hirogava/swifty-gasprom/backend/internal/errors/db"
	gameErrors "github.com/Hirogava/swifty-gasprom/backend/internal/errors/game"
	gameModels "github.com/Hirogava/swifty-gasprom/backend/internal/models/game"
	gameService "github.com/Hirogava/swifty-gasprom/backend/internal/service/game"
)

func (manager *Manager) SaveGame(game *gameModels.UserProgress) error {
	err := manager.Conn.QueryRow(`INSERT INTO user_progress (user_id, money, happiness, month, status) VALUES ($1, $2, $3, $4, $5)`, game.UserID, game.Money, game.Happiness, game.Month, game.Status).Scan(&game.ID)
	if err != nil {
		return err
	}

	return nil
}

func (manager *Manager) InitStocksAndCryptos(userID string) error {
	tx, err := manager.Conn.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	rows, err := tx.Query(`SELECT id, title, min_price, start_win_chance, type, crypto_category FROM scenarios where type = $1 or type = $2`, gameModels.Stocks, gameModels.Crypto)
	if err != nil {
		if err == sql.ErrNoRows {
			return dbErrors.ErrNotFound
		} else {
			return err
		}
	}

	for rows.Next() {
		var crypto gameModels.Risk
		var minPrice float64

		if err := rows.Scan(&crypto.ID, &crypto.Title, &minPrice, &crypto.WinChance, &crypto.RiskType, &crypto.CryptoCategory); err != nil {
			return err
		}

		rng := gameService.NewRNG()

		crypto.Price = rng.GenerateMinPrice(minPrice)

		if _, err := tx.Exec(`INSERT INTO user_crypto_scenarios (id, name, price, win_chance, type, crypto_category, month) VALUES ($1, $2, $3, $4, $5, $6, 1)`, crypto.ID, crypto.Title, crypto.Price, crypto.WinChance, crypto.RiskType, crypto.CryptoCategory); err != nil {
			return err
		}
	}
	rows.Close()

	rows, err = tx.Query(`SELECT id, title, min_price, start_win_chance, type FROM scenarios where type = $1`, gameModels.Bets)
	if err != nil {
		if err == sql.ErrNoRows {
			return dbErrors.ErrNotFound
		} else {
			return err
		}
	}
	defer rows.Close()

	for rows.Next() {
		var crypto gameModels.Risk

		if err := rows.Scan(&crypto.ID, &crypto.Title, &crypto.Price, &crypto.WinChance, &crypto.RiskType); err != nil {
			return err
		}

		if _, err := tx.Exec(`INSERT INTO user_crypto_scenarios (id, name, price, win_chance, type) VALUES ($1, $2, $3, $4, $5)`, crypto.ID, crypto.Title, crypto.Price, crypto.WinChance, crypto.RiskType); err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (manager *Manager) GetUserGameInfo(userID string) (*gameModels.UserGameInfo, error) {
	var gameInfo gameModels.UserGameInfo

	err := manager.Conn.QueryRow(`
		SELECT 
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
			uc.career_grade,
			uc.salary,
			c.name AS career_name,
			c.career_type
		FROM user_progress up
		LEFT JOIN current_progress_news cpn 
			ON up.user_id = cpn.user_id
		LEFT JOIN news n 
			ON cpn.news_id = n.id
		LEFT JOIN user_career uc 
			ON up.user_id = uc.user_id
		LEFT JOIN careers c 
			ON uc.career_id = c.id
		WHERE up.user_id = $1
		ORDER BY uc.taked_at DESC
		LIMIT 1
		`).Scan(&gameInfo.UserProgress.ID, &gameInfo.UserProgress.UserID, &gameInfo.UserProgress.Money,
		&gameInfo.UserProgress.Happiness, &gameInfo.UserProgress.Month, &gameInfo.UserProgress.Status,
		&gameInfo.News.ID, &gameInfo.News.Title, &gameInfo.News.Text, &gameInfo.Career.ID,
		&gameInfo.Career.Grade, &gameInfo.Career.Salary, &gameInfo.Career.Name, &gameInfo.Career.CareerType)
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
	defer rows.Close()

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

func (manager *Manager) GetVacancies() ([]gameModels.Vacancy, error) {
	var vacancies []gameModels.Vacancy

	rows, err := manager.Conn.Query(`SELECT id, name, career_type, career_level, career_grade, min_salary, max_salary, happiness_penalty_factor FROM careers WHERE career_grade = 1`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var vacancy gameModels.Vacancy
		var min, max float64

		if err := rows.Scan(&vacancy.ID, &vacancy.Name, &vacancy.Type, &vacancy.CareerLevel, &vacancy.CareerGrade, &min, &max, &vacancy.HappinessEffect); err != nil {
			return nil, err
		}

		vacancy.Salary = gameService.ArithmeticMeanSalary(min, max)

		vacancies = append(vacancies, vacancy)
	}

	return vacancies, nil
}

func (manager *Manager) GetVacancy(name string) ([]gameModels.Vacancy, error) {
	var vacancies []gameModels.Vacancy

	rows, err := manager.Conn.Query(`SELECT id, name, career_type, career_level, career_grade, min_salary, max_salary, happiness_penalty_factor, months_to_grade FROM careers WHERE name = $1`, name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, dbErrors.ErrNotFound
		} else {
			return nil, err
		}
	}
	defer rows.Close()

	for rows.Next() {
		var vacancy gameModels.Vacancy
		var min, max float64

		if err := rows.Scan(&vacancy.ID, &vacancy.Name, &vacancy.Type, &vacancy.CareerLevel, &vacancy.CareerGrade, &min, &max, &vacancy.HappinessEffect, &vacancy.MonthsToGrade); err != nil {
			return nil, err
		}

		vacancy.Salary = gameService.ArithmeticMeanSalary(min, max)

		vacancies = append(vacancies, vacancy)
	}

	return vacancies, nil
}

func (manager *Manager) GetPlayerVacancy(userID string) (*gameModels.Vacancy, error) {
	var vacancy gameModels.Vacancy

	if err := manager.Conn.QueryRow(`
			SELECT 
				uc.career_level,
				uc.career_grade,
				uc.salary,
				c.name AS career_name,
				c.career_type
				c.happiness_penalty_factor
				c.months_to_grade
			FROM user_career uc
			JOIN careers c 
				ON uc.career_id = c.id
			WHERE uc.user_id = $1
			ORDER BY uc.taked_at DESC
			LIMIT 1
		`, userID).Scan(&vacancy.CareerLevel, &vacancy.CareerGrade, &vacancy.Salary, &vacancy.Name, &vacancy.Type, &vacancy.HappinessEffect, &vacancy.MonthsToGrade); err != nil {
		if err == sql.ErrNoRows {
			return nil, dbErrors.ErrNotFound
		} else {
			return nil, err
		}
	}

	return &vacancy, nil
}

func (manager *Manager) SetPlayerVacancy(userID string, req *gameModels.SetVacancyRequest) error {
	_, err := manager.Conn.Exec(`
		UPDATE user_career
		SET career_id = $1,
			salary = $2,
			user_id = $3,
			career_level = $4,
			career_grade = $5
		`, req.ID, gameService.ArithmeticMeanSalary(req.MinSalary, req.MaxSalary), userID, req.Level, req.Grade)
	if err != nil {
		return err
	}

	return nil
}

func (manager *Manager) DeletePlayerVacancy(userID string, vacancyID int) error {
	if _, err := manager.Conn.Exec(`DELETE FROM user_career WHERE career_id = $1 and user_id = $2`, vacancyID, userID); err != nil {
		return err
	}

	return nil
}

func (manager *Manager) UpdatePlayerVacancy(userID string, name string, oldVacancy *gameModels.UpdatePlayerVacancyRequest) (*gameModels.Vacancy, error) {
	var vacancy gameModels.Vacancy
	var min, max float64
	var newType gameModels.CareerType
	var level int

	err := manager.Conn.QueryRow("SELECT career_type, career_level FROM careers WHERE name = $1 LIMIT 1", name).Scan(&newType, &level)
	if err != nil {
		return nil, err
	}
	if oldVacancy.Type != newType {
		vacancy.CareerLevel = 1
		vacancy.CareerGrade = 1
	} else if (vacancy.CareerGrade - 2) < 1 {
		vacancy.CareerGrade = 1
		vacancy.CareerLevel = level
	} else {
		vacancy.CareerGrade = oldVacancy.Grade - 2
		vacancy.CareerLevel = level
	}

	err = manager.Conn.QueryRow(`SELECT id, min_salary, max_salary WHERE career_level = $1 and career_grade = $2 and name = $3`, 
		level, vacancy.CareerGrade, name).Scan(&vacancy.ID, &min, &max)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, dbErrors.ErrNotFound
		}

		return nil, err
	}

	vacancy.Salary = gameService.ArithmeticMeanSalary(min, max)

	_, err = manager.Conn.Exec(`
		UPDATE user_career
		SET career_id = $1,
			salary = $2,
			career_level = $3,
			career_grade = $4
		WHERE user_id = $5 and career_id = $6
		`, vacancy.ID, vacancy.Salary, vacancy.CareerLevel, vacancy.CareerGrade, userID, oldVacancy.ID)
	if err != nil {
		return nil, err
	}

	return &vacancy, err
}

func (manager *Manager) UpdatePlayerVacancyGrade(userID string, oldCareerId int) (float64, int, error) {
	var min, max float64
	var penalty, id int

	err := manager.Conn.QueryRow(`
		SELECT 
			c_next.min_salary,
			c_next.max_salary,
			c_next.happiness_penalty_factor,
			c_next.id
		FROM user_career uc
		JOIN careers c_curr 
			ON uc.career_id = c_curr.id
		JOIN careers c_next 
			ON c_next.career_type = c_curr.career_type
		AND c_next.career_level = uc.career_level
		AND c_next.career_grade = uc.career_grade + 1
		WHERE uc.user_id = $1
		`, userID).Scan(&min, &max, &penalty, &id)
	if err != nil {
		return 0, 0, err
	}

	_, err = manager.Conn.Exec(`
		UPDATE user_career
		SET career_id = $1,
			salary = $2,
			career_grade = career_grade + 1
		WHERE user_id = $3 and career_id = $4
		`, id, gameService.ArithmeticMeanSalary(min, max), userID, oldCareerId)

	return gameService.ArithmeticMeanSalary(min, max), penalty, err
}

func (manager *Manager) BuyBetsItem(req *gameModels.BetsRequest, userID string) error {
	tx, err := manager.Conn.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(`
		UPDATE user_progress
		SET money = money - $1
	`, req.Price)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
		INSERT INTO user_buyed_risks (
			user_id,
			risk_id
		)
	`)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (manager *Manager) BuyCryptoItem(req *gameModels.CryptoRequest, userID string) error {
	tx, err := manager.Conn.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(`
		UPDATE user_progress
		SET money = money - $1
	`, req.Price)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
		INSERT INTO user_buyed_risks (
			user_id,
			crypto_id
		)
	`)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (manager *Manager) BuyStocksItem(req *gameModels.StocksRequest, userID string) error {
	tx, err := manager.Conn.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(`
		UPDATE user_progress
		SET money = money - $1
	`, req.Price)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
		INSERT INTO user_buyed_risks (
			user_id,
			crypto_id
		)
	`)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (manager *Manager) BuyQPItem(req *gameModels.QuestionableProjectsRequest, userID string) error {
	tx, err := manager.Conn.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(`
		UPDATE user_progress
		SET money = money - $1
	`, req.Price)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
		INSERT INTO user_buyed_risks (
			user_id,
			risk_id
		)
	`)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (manager *Manager) GetRiskItems(userID string) (*gameModels.RiskItems, error) {
	var items gameModels.RiskItems

	rows, err := manager.Conn.Query(`
		SELECT name, type, price
		FROM risk
		WHERE user_id = $1
		`, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, dbErrors.ErrNotFound
		} else {
			return nil, err
		}
	}

	for rows.Next() {
		var item gameModels.Risk

		if err := rows.Scan(&item.Name, &item.RiskType, &item.Price); err != nil {
			return nil, err
		}

		switch item.RiskType {
		case gameModels.Bets:
			items.Bets = append(items.Bets, item)
		case gameModels.QuestionableProjects:
			items.QuestionableProjects = append(items.QuestionableProjects, item)
		}
	}
	rows.Close()

	rows, err = manager.Conn.Query(`
		SELECT name, type, price
		FROM user_crypto_scenarios
		WHERE user_id = $1
		AND month = (
			SELECT MAX(month)
			FROM user_crypto_scenarios
			WHERE user_id = $1
		)`, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, dbErrors.ErrNotFound
		} else {
			return nil, err
		}
	}

	for rows.Next() {
		var item gameModels.Risk

		if err := rows.Scan(&item.Name, &item.RiskType, &item.Price); err != nil {
			return nil, err
		}

		switch item.RiskType {
		case gameModels.Crypto:
			items.Crypto = append(items.Crypto, item)
		case gameModels.Stocks:
			items.Stocks = append(items.Stocks, item)
		}
	}

	return &items, nil
}

func (manager *Manager) GetPlayerRiskItems(userID string) (*gameModels.RiskItems, error) {
	var items gameModels.RiskItems

	rows, err := manager.Conn.Query(`
		SELECT 
			COALESCE(r.name, ucs.name) AS name,
			COALESCE(r.type, ucs.type) AS type,
			COALESCE(r.price, ucs.price) AS price
		FROM user_buyed_risks ubr
		LEFT JOIN risk r 
			ON ubr.risk_id = r.id
		LEFT JOIN user_crypto_scenarios ucs 
			ON ubr.crypto_id = ucs.id
		WHERE ubr.user_id = $1
		`, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, dbErrors.ErrNotFound
		} else {
			return nil, err
		}
	}

	for rows.Next() {
		var item gameModels.Risk

		if err := rows.Scan(&item.Name, &item.RiskType, &item.Price); err != nil {
			return nil, err
		}

		switch item.RiskType {
		case gameModels.Crypto:
			items.Crypto = append(items.Crypto, item)
		case gameModels.Stocks:
			items.Stocks = append(items.Stocks, item)
		case gameModels.Bets:
			items.Bets = append(items.Bets, item)
		case gameModels.QuestionableProjects:
			items.QuestionableProjects = append(items.QuestionableProjects, item)
		}
	}

	return &items, nil
}

func (manager *Manager) GetRiskItem(itemId int, userID string, itemType string) (*gameModels.RiskItem, error) {
	var item gameModels.RiskItem
	var history []byte

	err := manager.Conn.QueryRow(`
		SELECT 
			ucs.id,
			ucs.name,
			CASE 
				WHEN $3 IN ('crypto', 'stocks') THEN 
					json_agg(
						json_build_object(
							'price', ucs.price,
							'month', ucs.month
						) ORDER BY ucs.month
					)
				ELSE NULL
			END AS price_history,
			(
				SELECT price
				FROM user_crypto_scenarios u2
				WHERE u2.id = ucs.id
				ORDER BY u2.month DESC
				LIMIT 1
			) AS current_price,
			MAX(ucs.win_chance) AS win_chance,
			MAX(ucs.lose_chance) AS lose_chance
		FROM user_crypto_scenarios ucs
		WHERE ucs.id = $1 AND ucs.user_id = $2
		GROUP BY ucs.id, ucs.name
	`, itemId, userID, itemType).Scan(&item.ID, &item.Name, &history, &item.CurrentPrice, &item.WinChance, &item.LoseChance)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, dbErrors.ErrNotFound
		}
		return nil, err
	}

	if itemType == string(gameModels.Crypto) || itemType == string(gameModels.Stocks) {
		if len(history) > 0 {
			if err := json.Unmarshal(history, &item.PriceHistory); err != nil {
				return nil, err
			}
		}
	}

	return &item, nil
}

func (manager *Manager) GetRiskItemsByType(itemType string, userID string) ([]gameModels.Risk, error) {
	var items []gameModels.Risk

	if itemType == string(gameModels.Crypto) || itemType == string(gameModels.Stocks) {
		rows, err := manager.Conn.Query(`
			SELECT 
				name,
				price,
				win_chance,
				lose_chance
			FROM user_crypto_scenarios
			WHERE type = $1 and user_id = $2
			`, itemType, userID)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, dbErrors.ErrNotFound
			} else {
				return nil, err
			}
		}
		defer rows.Close()

		for rows.Next() {
			var item gameModels.Risk

			if err := rows.Scan(&item.Name, &item.Price, &item.WinChance, &item.LoseChance); err != nil {
				return nil, err
			}

			items = append(items, item)
		}
	}
	
	rows, err := manager.Conn.Query(`
		SELECT 
			name,
			price,
			win_chance,
			lose_chance
		FROM risk
		WHERE type = $1 and user_id = $2
		`, itemType, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, dbErrors.ErrNotFound
		} else {
			return nil, err
		}
	}
	defer rows.Close()

	for rows.Next() {
		var item gameModels.Risk

		if err := rows.Scan(&item.Name, &item.Price, &item.WinChance, &item.LoseChance); err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
}

func (manager *Manager) GetPlayerNews(userID string) ([]gameModels.News, error) {
	var news []gameModels.News

	rows, err := manager.Conn.Query(`
		SELECT 
			n.id,
			n.news_title,
			n.news_text
		FROM current_progress_news cpn
		JOIN news n ON n.id = cpn.news_id
		WHERE cpn.user_id = $1
		AND cpn.month = (
			SELECT MAX(month) 
			FROM current_progress_news 
			WHERE user_id = $1
		)
		ORDER BY n.id
		`, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, dbErrors.ErrNotFound
		} else {
			return nil, err
		}
	}
	defer rows.Close()

	for rows.Next() {
		var item gameModels.News

		if err := rows.Scan(&item.ID, &item.Title, &item.Text); err != nil {
			return nil, err
		}

		news = append(news, item)
	}

	return news, nil
}

func (manager *Manager) GetCurrentNews(newsID int, userID string) (*gameModels.News, error) {
	var news gameModels.News

	if err := manager.Conn.QueryRow(`
			SELECT 
				n.id,
				n.type,
				n.news_title,
				n.news_text
			FROM current_progress_news cpn
			JOIN news n ON n.id = cpn.news_id
			WHERE cpn.user_id = $1
			AND cpn.news_id = $2
			AND cpn.month = (
				SELECT MAX(month) 
				FROM current_progress_news 
				WHERE user_id = $1
			)
			`, userID, newsID).Scan(&news.ID, &news.Type, &news.Title, &news.Text); err != nil {
		if err == sql.ErrNoRows {
			return nil, dbErrors.ErrNotFound
		} else {
			return nil, err
		}
	}

	return &news, nil
}
