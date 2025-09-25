package postgres

import (
	"database/sql"

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

	rows, err := manager.Conn.Query(`SELECT id, name, career_type, career_level, career_grade, min_salary, max_salary, happiness_penalty_factor FROM careers WHERE name = $1`, name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, dbErrors.ErrNotFound
		} else {
			return nil, err
		}
	}

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

func (manager *Manager) GetPlayerVacancy(userID string) (*gameModels.Vacancy, error) {
	var vacancy gameModels.Vacancy

	if err := manager.Conn.QueryRow(`SELECT 
				uc.career_level,
				uc.career_grade,
				uc.salary,
				c.name AS career_name,
				c.career_type
				c.happiness_penalty_factor
			FROM user_career uc
			JOIN careers c 
				ON uc.career_id = c.id
			WHERE uc.user_id = $1
			ORDER BY uc.taked_at DESC
			LIMIT 1
		`, userID).Scan(&vacancy.CareerLevel, &vacancy.CareerGrade, &vacancy.Salary, &vacancy.Name, &vacancy.Type, &vacancy.HappinessEffect); err != nil {
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
