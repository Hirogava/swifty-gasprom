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

type CareerType string

const (
	ITAndTechnology CareerType = "it_and_technology"
	EngineeringAndManufacturing CareerType = "engineering_and_manufacturing"
	MedicineAndHealthcare CareerType = "medicine_and_healthcare"
	MarketingAndSales CareerType = "marketing_and_sales"
	WorkingProfessions CareerType = "working_professions"
)

type RiskType string

const (
	Bets RiskType = "bets"
	QuestionableProjects RiskType = "questionable_projects"
	Crypto RiskType = "crypto"
	Stocks RiskType = "stocks"
)

type NewsType string

const (
	Economic NewsType = "economic"
	Political NewsType = "political"
	Corporate NewsType = "corporate"
	Useless NewsType = "useless"
)
