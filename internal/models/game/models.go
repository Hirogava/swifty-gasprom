package game

type UserProgress struct {
	ID int64 `json:"id"`
	UserID string `json:"user_id"`
	Money float64 `json:"money"`
	Happiness int `json:"happiness"`
	Month int `json:"month"`
	Status UserStatus `json:"status"`
}

type News struct {
	ID int64 `json:"id"`
	Title string `json:"title"`
	Text string `json:"text"`
}

type Career struct {
	ID int64 `json:"id"`
	Level int `json:"level"`
	Name string `json:"name"`
	CareerField string `json:"career_field"`
}

type UserGameInfo struct {
	UserProgress UserProgress
	News News
	Career Career
}

type Event struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	EventText string `json:"event_text"`
	CapitalPercent int `json:"capital_percent"`
	ToAgree int `json:"to_agree"`
	Refuse int `json:"refuse"`
}

type LifeMarketItem struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	Cost float64 `json:"cost"`
	Category LifeMarketCategory `json:"category"`
	HappinessEffect int `json:"happiness_effect"`
}
