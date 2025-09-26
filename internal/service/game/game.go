package game

import (
	"encoding/json"
	"math/rand"
	"time"

	gameErrors "github.com/Hirogava/swifty-gasprom/backend/internal/errors/game"
	gameModels "github.com/Hirogava/swifty-gasprom/backend/internal/models/game"
)

func StartGame(userID string) *gameModels.UserProgress {
	var progress gameModels.UserProgress

	progress.Happiness = 100
	progress.Money = 0
	progress.Month = 1
	progress.Status = gameModels.UserStatusActive
	progress.UserID = userID

	return &progress
}

func ArithmeticMeanSalary(minSalary, maxSalary float64) float64 {
	if minSalary == maxSalary {
		return minSalary
	} else {
		return (minSalary + maxSalary) / 2
	}
}

type Outcome string

const (
	Win  Outcome = "win"
	Loss Outcome = "loss"
)

type RNG struct {
	r *rand.Rand
}

func NewRNG() *RNG {
	src := rand.NewSource(time.Now().UnixNano())
	return &RNG{r: rand.New(src)}
}

func (rng *RNG) GetOutcome(winChance float64) Outcome {
	r := rng.r.Float64()

	if r < winChance {
		return Win
	}

	return Loss
}

func (rng *RNG) GenerateMinPrice(price float64) float64 {
	return price + rng.r.Float64() * 500
}

func EnvelopeRiskStruct(envelope *gameModels.Envelope) (any, error) {
	msgStruct := getMessageType(envelope.Type)
	if msgStruct == nil {
		return nil, gameErrors.ErrInvalidMessageType
	}

	err := json.Unmarshal(envelope.Data, msgStruct)
	if err != nil {
		return nil, err
	}

	return msgStruct, nil
}

func getMessageType(msgType gameModels.RiskType) any {
	switch msgType {
	case gameModels.Crypto:
		return &gameModels.CryptoRequest{}
	case gameModels.Bets:
		return &gameModels.BetsRequest{}
	case gameModels.Stocks:
		return &gameModels.StocksRequest{}
	case gameModels.QuestionableProjects:
		return &gameModels.QuestionableProjectsRequest{}
	default:
		return nil
	}
}
