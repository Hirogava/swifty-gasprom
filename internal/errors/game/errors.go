package game

import "errors"

var (
	ErrNotEnoughMoney = errors.New("not enough money")
	ErrPlayerVacancyNotFound = errors.New("player vacancy not found")
	ErrInsufficientPlayerLevel = errors.New("the player's level must be lower by a maximum of one job level.")
	ErrInvalidMessageType = errors.New("msg type is invalid")
)
