package bank

type CardType string

const (
	DebitCard CardType = "debit_card"
	CreditCard CardType = "credit_card"
)

type BankBonusType string

const (
	Cashback BankBonusType = "cashback"
	Discount BankBonusType = "discount"
)

type BankBonusStatus string

const (
	Claimed BankBonusStatus = "claimed"
	Pending BankBonusStatus = "pending"
	Expired BankBonusStatus = "expired"
)