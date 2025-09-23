package bank

type BankProduct struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Url string `json:"url"`
	Type CardType `json:"type"`
}

type BankBonus struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Url string `json:"url"`
	Type BankBonusType `json:"type"`
}
