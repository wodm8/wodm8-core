package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/huandu/go-sqlbuilder"
	"github.com/wodm8/wodm8-core/internal/crossfit"
	"time"
)

type WodRoundRepository struct {
	db        *sql.DB
	dbTimeout time.Duration
}

func NewWodRoundRepository(db *sql.DB, dbTimeout time.Duration) *WodRoundRepository {
	return &WodRoundRepository{
		db:        db,
		dbTimeout: dbTimeout,
	}
}

func (r *WodRoundRepository) Save(ctx context.Context, wodRound crossfit.WodRound) error {
	wodRoundSqlStruct := sqlbuilder.NewStruct(new(sqlWodRound))
	query, args := wodRoundSqlStruct.InsertInto(sqlWodRoundTable, sqlWodRound{
		ID:                 wodRound.ID().String(),
		WodId:              wodRound.WODId().String(),
		SetNumber:          wodRound.SetNumber().Int(),
		RoundNumber:        wodRound.RoundNumber().Int(),
		RepetitionsByRound: wodRound.RepetitionsByRound().Int(),
		Time:               wodRound.Time().Int(),
		RestTime:           wodRound.RestTime().Int(),
		RemainingRest:      wodRound.RemainingTime().Bool(),
	}).Build()

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	_, err := r.db.ExecContext(ctxTimeout, query, args...)
	if err != nil {
		fmt.Printf("error saving wod round: %v", err)
		return err
	}
	return nil
}
