package game

type UserStatus string

const (
	UserStatusActive UserStatus = "active"
	UserStatusFinished UserStatus = "finished"
	UserStatusBurnout UserStatus = "burnout"
	UserStatusBankrupt UserStatus = "bankrupt"
)