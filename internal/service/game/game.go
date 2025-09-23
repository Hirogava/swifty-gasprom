package game

import gameModels "github.com/Hirogava/swifty-gasprom/backend/internal/models/game"

func StartGame(userID string) *gameModels.UserProgress {
	var progress gameModels.UserProgress

	progress.Happiness = 100
	progress.Money = 0
	progress.Month = 1
	progress.Status = gameModels.UserStatusActive
	progress.UserID = userID

	return &progress
}