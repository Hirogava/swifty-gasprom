package postgres

import (
	"database/sql"

	dbErrors "github.com/Hirogava/swifty-gasprom/backend/internal/errors/db"
	bankModels "github.com/Hirogava/swifty-gasprom/backend/internal/models/bank"
)

func (manager *Manager) SaveUserBankProduct(userId string, cardId int) (bankModels.BankProduct, error) {
	var bankProduct bankModels.BankProduct

	if _, err := manager.Conn.Exec(`INSERT INTO user_used_bank_products (user_id, bank_product_id) VALUES ($1, $2)`, userId, cardId); err != nil {
		return bankModels.BankProduct{}, err
	}

	if err := manager.Conn.QueryRow(`SELECT name, url, type FROM bank_products WHERE id = $1`, cardId).Scan(&bankProduct.Name, &bankProduct.Url, &bankProduct.Type); err != nil {
		if err == sql.ErrNoRows {
			return bankModels.BankProduct{}, dbErrors.ErrNotFound
		} else {
			return bankModels.BankProduct{}, err
		}
	}

	return bankProduct, nil
}

func (manager *Manager) GetUserBankProduct(userId string) (bankModels.BankProduct, error) {
	var bankProduct bankModels.BankProduct

	if err := manager.Conn.QueryRow(`SELECT 
				bp.name,
				bp.url,
				bp.type
			FROM user_used_bank_products uubp
			JOIN bank_products bp 
				ON uubp.bank_product_id = bp.id
			WHERE uubp.user_id = $1`,
		userId).Scan(&bankProduct.Name, &bankProduct.Url, &bankProduct.Type); err != nil {
		if err == sql.ErrNoRows {
			return bankModels.BankProduct{}, dbErrors.ErrNotFound
		} else {
			return bankModels.BankProduct{}, err
		}
	}
	
	return bankProduct, nil
}

func (manager *Manager) SaveUserBankBonus(userId string, bonusId int) (bankModels.BankBonus, error) {
	var bankBonus bankModels.BankBonus

	if _, err := manager.Conn.Exec(`INSERT INTO user_bonuses (user_id, bonus_id, status) VALUES ($1, $2, $3)`, userId, bonusId, bankModels.Pending); err != nil {
		return bankModels.BankBonus{}, err
	}

	if err := manager.Conn.QueryRow(`SELECT name, url, type FROM bank_bonuses WHERE id = $1`, bonusId).Scan(&bankBonus.Name, &bankBonus.Url, &bankBonus.Type); err != nil {
		if err == sql.ErrNoRows {
			return bankModels.BankBonus{}, dbErrors.ErrNotFound
		} else {
			return bankModels.BankBonus{}, err
		}
	}

	return bankBonus, nil
}

func (manager *Manager) GetUserBankBonuses(userId string) ([]bankModels.BankBonus, error) {
	var bankBonuses []bankModels.BankBonus

	rows, err := manager.Conn.Query(`SELECT 
				bb.name,
				bb.url,
				bb.type
				FROM user_bonuses ub
				JOIN bank_bonuses bb
				ON ub.bonus_id = bb.id
				WHERE ub.user_id = $1`, userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return []bankModels.BankBonus{}, dbErrors.ErrNotFound
		} else {
			return []bankModels.BankBonus{}, err
		}
	}
	defer rows.Close()

	for rows.Next() {
		var bankBonus bankModels.BankBonus

		if err := rows.Scan(&bankBonus.Name, &bankBonus.Url, &bankBonus.Type); err != nil {
			return []bankModels.BankBonus{}, err
		}

		bankBonuses = append(bankBonuses, bankBonus)
	}

	return bankBonuses, nil
}

func (manager *Manager) UpdateUserBonusStatusType(req bankModels.UpdateBonusStatusRequest) error {
	if _, err := manager.Conn.Exec(`UPDATE user_bonuses SET status = $1 WHERE bonus_id = $2 and user_id = $3`, req.Status, req.ID, req.UserID); err != nil {
		return err
	}

	return nil
}
