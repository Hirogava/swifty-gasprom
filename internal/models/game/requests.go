package game

type LifeMarketItemRequest struct {
	ID        int64   `json:"id" binding:"required"`
	Cost      float64 `json:"cost" binding:"required"`
	Happiness int     `json:"happiness" binding:"required"`
}

type GetVacancyRequest struct {
	Name string `json:"name" binding:"required"`
}

type SetVacancyRequest struct {
	ID        int64   `json:"id" binding:"required"`
	Level     int     `json:"level" binding:"required"`
	Grade     int     `json:"grade" binding:"required"`
	MinSalary float64 `json:"min_salary" binding:"required"`
	MaxSalary float64 `json:"max_salary" binding:"required"`
}

type UpdatePlayerVacancyRequest struct {
	ID    int64      `json:"id" binding:"required"`
	Level int        `json:"level" binding:"required"`
	Grade int        `json:"grade" binding:"required"`
	Type  CareerType `json:"type" binding:"required"`
}

type Envelope struct {
	Type RiskType `json:"type" binding:"required"`
	Data []byte `json:"data" binding:"omitempty"`
}

type CryptoRequest struct {
	Price float64 `json:"price" binding:"required"`
	CryptoID int64 `json:"crypto_id" binding:"required"`
}

type StocksRequest struct {
	Price float64 `json:"price" binding:"required"`
	StockID int64 `json:"stock_id" binding:"required"`
}

type BetsRequest struct {
	Price float64 `json:"price" binding:"required"`
	BetID int64 `json:"bet_id" binding:"required"`
}

type QuestionableProjectsRequest struct {
	Price float64 `json:"price" binding:"required"`
	QPID int64 `json:"qp_id" binding:"required"`
}
