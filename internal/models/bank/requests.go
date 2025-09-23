package bank

type UpdateBonusStatusRequest struct {
	ID int `json:"id" binding:"required"`
	UserID string `json:"user_id" binding:"required"`
	Status BankBonusStatus `json:"status" binding:"required"`
}