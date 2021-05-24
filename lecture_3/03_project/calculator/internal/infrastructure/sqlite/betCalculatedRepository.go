package sqlite

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"

	domainmodels "github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/domain/models"
	storagemodels "github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/infrastructure/sqlite/models"
)

// BetCalculatedRepository provides methods that operate on calculated bets SQLite database.
type BetCalculatedRepository struct {
	dbExecutor          DatabaseExecutor
	betCalculatedMapper BetCalculatedMapper
}

// NewBetCalculatedRepository creates and returns a new BetCalculatedRepository.
func NewBetCalculatedRepository(dbExecutor DatabaseExecutor, betCalculatedMapper BetCalculatedMapper) *BetCalculatedRepository {
	return &BetCalculatedRepository{
		dbExecutor:          dbExecutor,
		betCalculatedMapper: betCalculatedMapper,
	}
}

// InsertBetCalculated inserts the provided calculated bet into the database. An error is returned if the operation
// has failed.
func (r *BetCalculatedRepository) InsertBetCalculated(ctx context.Context, bet domainmodels.BetCalculated) error {
	storageBet := r.betCalculatedMapper.MapDomainBetToStorageBet(bet)
	err := r.queryInsertBetCalculated(ctx, storageBet)
	if err != nil {
		return errors.Wrap(err, "calculated bet repository failed to insert a calculated bet with id "+bet.Id)
	}
	return nil
}

func (r *BetCalculatedRepository) queryInsertBetCalculated(ctx context.Context, bet storagemodels.BetCalculated) error {
	insertBetSQL := "INSERT INTO bets(id, selection_id, selection_coefficient, payment) VALUES (?, ?, ?, ?)"
	statement, err := r.dbExecutor.PrepareContext(ctx, insertBetSQL)
	if err != nil {
		return err
	}

	_, err = statement.ExecContext(ctx, bet.Id, bet.SelectionId, bet.SelectionCoefficient, bet.Payment)
	return err
}

// UpdateBetCalculated updates the provided bet in the database. An error is returned if the operation
// has failed.
func (r *BetCalculatedRepository) UpdateBetCalculated(ctx context.Context, bet domainmodels.BetCalculated) error {
	storageBet := r.betCalculatedMapper.MapDomainBetToStorageBet(bet)
	err := r.queryUpdateBetCalculated(ctx, storageBet)
	if err != nil {
		return errors.Wrap(err, "bet repository failed to update a calculated bet with id "+bet.Id)
	}
	return nil
}

func (r *BetCalculatedRepository) queryUpdateBetCalculated(ctx context.Context, bet storagemodels.BetCalculated) error {
	updateBetSQL := "UPDATE bets SET selection_id=?, selection_coefficient=?, payment=? WHERE id=?"

	statement, err := r.dbExecutor.PrepareContext(ctx, updateBetSQL)
	if err != nil {
		return err
	}

	_, err = statement.ExecContext(ctx, bet.SelectionId, bet.SelectionCoefficient, bet.Payment, bet.Id)
	return err
}

// GetBetCalculatedByID fetches a calculated bet from the database and returns it. The second returned value indicates
// whether the bet exists in DB. If the bet does not exist, an error will not be returned.
func (r *BetCalculatedRepository) GetBetCalculatedByID(ctx context.Context, id string) (domainmodels.BetCalculated, bool, error) {
	storageBet, err := r.queryGetBetCalculatedByID(ctx, id)
	if err == sql.ErrNoRows {
		return domainmodels.BetCalculated{}, false, nil
	}
	if err != nil {
		return domainmodels.BetCalculated{}, false, errors.Wrap(err, "bet repository failed to get a calculated bet with id "+id)
	}

	domainBet := r.betCalculatedMapper.MapStorageBetToDomainBet(storageBet)
	return domainBet, domainBet != domainmodels.BetCalculated{}, nil
}

func (r *BetCalculatedRepository) queryGetBetCalculatedByID(ctx context.Context, id string) (storagemodels.BetCalculated, error) {
	row, err := r.dbExecutor.QueryContext(ctx, "SELECT * FROM bets WHERE id='"+id+"';")
	if err != nil {
		return storagemodels.BetCalculated{}, err
	}
	defer row.Close()

	// This will move to the "next" result (which is the only result, because a single bet is fetched).
	hasNext := row.Next()
	if !hasNext {
		return storagemodels.BetCalculated{}, nil
	}

	var selectionId string
	var selectionCoefficient int
	var payment int

	err = row.Scan(&id, &selectionId, &selectionCoefficient, &payment)
	if err != nil {
		return storagemodels.BetCalculated{}, err
	}

	return storagemodels.BetCalculated{
		Id:                   id,
		SelectionId:          selectionId,
		SelectionCoefficient: selectionCoefficient,
		Payment:              payment,
	}, nil
}

func (r *BetCalculatedRepository) GetBetBySelectionID(ctx context.Context, selectionId string) ([]domainmodels.BetCalculated, bool, error) {
	storageBets, err := r.queryGetBetsBySelectionID(ctx, selectionId)
	if err == sql.ErrNoRows {
		return []domainmodels.BetCalculated{}, false, nil
	}
	if err != nil {
		return []domainmodels.BetCalculated{}, false, errors.Wrap(err, "bet repository failed to get a bets with selection id "+selectionId)
	}

	var domainBets []domainmodels.BetCalculated

	for _, bet := range storageBets {
		domainBet := r.betCalculatedMapper.MapStorageBetToDomainBet(bet)
		domainBets = append(domainBets, domainBet)
	}

	return domainBets, true, nil
}

func (r *BetCalculatedRepository) queryGetBetsBySelectionID(ctx context.Context, selectionId string) ([]storagemodels.BetCalculated, error) {
	row, err := r.dbExecutor.QueryContext(ctx, "SELECT * FROM bets WHERE selection_id='"+selectionId+"';")
	if err != nil {
		return []storagemodels.BetCalculated{}, err
	}
	defer row.Close()

	var calculatedBets []storagemodels.BetCalculated

	found := row.Next()
	for found {
		var id string
		var selectionCoefficient int
		var payment int

		err = row.Scan(&id, &selectionId, &selectionCoefficient, &payment)
		if err != nil {
			found = row.Next()
			continue
		}

		calculatedBets = append(calculatedBets, storagemodels.BetCalculated{
			Id:                   id,
			SelectionId:          selectionId,
			SelectionCoefficient: selectionCoefficient,
			Payment:              payment,
		})
		found = row.Next()
	}

	return calculatedBets, nil
}
