package game

import "database/sql"

type UserProgress struct {
	ID               int64      `json:"id"`
	UserID           string     `json:"user_id"`
	Money            float64    `json:"money"`
	Happiness        int        `json:"happiness"`
	Month            int        `json:"month"`
	Status           UserStatus `json:"status"`
	Natural_expenses *int        `json:"natural_expenses"`
}

type News struct {
	ID    *int64    `json:"id"`
	Title *string   `json:"title"`
	Text  *string   `json:"text"`
	Type  *NewsType `json:"type"`
}

type Career struct {
	ID         *int64      `json:"id"`
	Grade      *int        `json:"grade"`
	Salary     *float64    `json:"salary"`
	Name       *string     `json:"name"`
	CareerType *CareerType `json:"career_type"`
}

type UserGameInfo struct {
	UserProgress UserProgress
	News         News
	Career       Career
}

type Event struct {
	ID             int64  `json:"id"`
	Name           string `json:"name"`
	EventText      string `json:"event_text"`
	CapitalPercent int    `json:"capital_percent"`
	ToAgree        int    `json:"to_agree"`
	Refuse         int    `json:"refuse"`
}

type LifeMarketItem struct {
	ID              int64              `json:"id"`
	Name            string             `json:"name"`
	Cost            float64            `json:"cost"`
	Category        LifeMarketCategory `json:"category"`
	HappinessEffect int                `json:"happiness_effect"`
}

type Vacancy struct {
	ID              int64      `json:"id"`
	Name            string     `json:"name"`
	Type            CareerType `json:"type"`
	CareerLevel     int        `json:"career_level"`
	CareerGrade     int        `json:"career_grade"`
	Salary          float64    `json:"salary"`
	HappinessEffect int        `json:"happiness_effect"`
	MonthsToGrade   int        `json:"months_to_grade"`
}

type Risk struct {
	ID             int64         `json:"id"`
	Name           string        `json:"name"`
	Title          string        `json:"title"`
	RiskType       RiskType      `json:"risk_type"`
	CryptoCategory int           `json:"crypto_category"`
	Price          float64       `json:"price"`
	WinChance      int           `json:"win_chance"`
	LoseChance     sql.NullInt64 `json:"lose_chance"`
}

type RiskItem struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	PriceHistory []struct {
		Price float64 `json:"price"`
		Month int     `json:"month"`
	} `json:"price_history"`
	CurrentPrice float64 `json:"current_price"`
	WinChance    int     `json:"win_chance"`
	LoseChance   int     `json:"lose_chance"`
}

type RiskItems struct {
	Crypto               []Risk `json:"crypto"`
	Stocks               []Risk `json:"stocks"`
	Bets                 []Risk `json:"bets"`
	QuestionableProjects []Risk `json:"questionable_projects"`
}

type DBRisk struct {
	ID             int64    `json:"id"`
	Name           string   `json:"name"`
	Title          string   `json:"title"`
	RiskType       RiskType `json:"risk_type"`
	CryptoCategory *int     `json:"crypto_category"`
	Price          float64  `json:"price"`
	MaxPrice       float64  `json:"max_price"`
	MinPrice       float64  `json:"min_price"`
	WinChance      int      `json:"win_chance"`
	LoseChance     *int      `json:"lose_chance"`
	MaxWin         float64  `json:"max_win"`
	MaxLose        float64  `json:"max_lose"`
}

type DBNews struct {
	ID             int64    `json:"id"`
	Title          string   `json:"title"`
	Text           string   `json:"text"`
	Type           NewsType `json:"type"`
	CryptoCategory int      `json:"crypto_category"`
	EffectOnMarket int      `json:"effect_on_market"`
	Effect         bool     `json:"effect"`
}

type Month struct {
	Money     float64 `json:"money"`
	Happiness int     `json:"happiness"`
	Month     int     `json:"month"`
	News      []News  `json:"news"`
}
