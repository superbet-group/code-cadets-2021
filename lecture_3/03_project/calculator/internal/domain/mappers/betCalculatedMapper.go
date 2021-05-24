package mappers

import (
	"math"

	domainmodels "github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/domain/models"
	storagemodels "github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/infrastructure/sqlite/models"
)

// BetCalculatedMapper maps storage calculated bets to domain calculated bets and vice versa.
type BetCalculatedMapper struct {
}

// NewBetCalculatedMapper creates and returns a new BetCalculatedMapper.
func NewBetCalculatedMapper() *BetCalculatedMapper {
	return &BetCalculatedMapper{}
}

// MapDomainBetToStorageBet maps the given domain calculated bet into storage calculated bet. Floating point values will
// be converted to corresponding integer values of the storage bet by multiplying them with 100.
func (m *BetCalculatedMapper) MapDomainBetToStorageBet(calculatedBet domainmodels.BetCalculated) storagemodels.BetCalculated {
	return storagemodels.BetCalculated{
		Id:                   calculatedBet.Id,
		SelectionId:          calculatedBet.SelectionId,
		SelectionCoefficient: int(math.Round(calculatedBet.SelectionCoefficient * 100)),
		Payment:              int(math.Round(calculatedBet.Payment * 100)),
	}
}

// MapStorageBetToDomainBet maps the given storage bet into domain bet. Floating point values will
// be converted from corresponding integer values of the storage bet by dividing them with 100.
func (m *BetCalculatedMapper) MapStorageBetToDomainBet(storageBet storagemodels.BetCalculated) domainmodels.BetCalculated {
	return domainmodels.BetCalculated{
		Id:                   storageBet.Id,
		SelectionId:          storageBet.SelectionId,
		SelectionCoefficient: float64(storageBet.SelectionCoefficient) / 100,
		Payment:              float64(storageBet.Payment) / 100,
	}
}
