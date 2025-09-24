package game

type UserStatus string

const (
	UserStatusActive UserStatus = "active"
	UserStatusFinished UserStatus = "finished"
	UserStatusBurnout UserStatus = "burnout"
	UserStatusBankrupt UserStatus = "bankrupt"
)

type LifeMarketCategory string

const (
	EntertainmentAndRecreation LifeMarketCategory = "entertainment_and_recreation"
	AppliancesAndGadgets LifeMarketCategory = "appliances_and_gadgets"
	GiftsAndSocialInteraction LifeMarketCategory = "gifts_and_social_interaction"
	EducationAndSelfDevelopment LifeMarketCategory = "education_and_self_development"
	HealthAndCare LifeMarketCategory = "health_and_care"
	EverydayJoys LifeMarketCategory = "everyday_joys"
)
