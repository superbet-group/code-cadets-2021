package handler

import (
	"context"

	domainmodels "github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/domain/models"
)

type BetCalculatedRepository interface {
	InsertBetCalculated(ctx context.Context, bet domainmodels.BetCalculated) error
	UpdateBetCalculated(ctx context.Context, bet domainmodels.BetCalculated) error
	GetBetCalculatedByID(ctx context.Context, id string) (domainmodels.BetCalculated, bool, error)
	GetBetBySelectionID(ctx context.Context, selectionId string) ([]domainmodels.BetCalculated, bool, error)
}
