package game

type LifeMarketItemRequest struct {
	ID int64 `json:"id" binding:"required"`
	Cost float64 `json:"cost" binding:"required"`
	Happiness int `json:"happiness" binding:"required"`
}

type GetVacancyRequest struct {
	Name string `json:"name" binding:"required"`
}

type SetVacancyRequest struct {
	ID int64 `json:"id" binding:"required"`
	Level int `json:"level" binding:"required"`
	Grade int `json:"grade" binding:"required"`
	MinSalary float64 `json:"min_salary" binding:"required"`
	MaxSalary float64 `json:"max_salary" binding:"required"`
}

type UpdatePlayerVacancyRequest struct {
	ID int64 `json:"id" binding:"required"`
	Level int `json:"level" binding:"required"`
	Grade int `json:"grade" binding:"required"`
	Type CareerType `json:"type" binding:"required"`
}
