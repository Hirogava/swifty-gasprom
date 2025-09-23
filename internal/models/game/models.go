package game

type UserProgress struct {
	ID int64 `json:"id"`
	UserID string `json:"user_id"`
	Money float64 `json:"money"`
	Happiness int `json:"happiness"`
	Month int `json:"month"`
	Status UserStatus `json:"status"`
}