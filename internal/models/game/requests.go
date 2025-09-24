package game

type LifeMarketItemRequest struct {
	ID int64 `json:"id" binding:"required"`
	Cost float64 `json:"cost" binding:"required"`
	Happiness int `json:"happiness" binding:"required"`
}