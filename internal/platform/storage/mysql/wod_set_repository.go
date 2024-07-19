package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/wodm8/wodm8-core/internal/crossfit"

	"github.com/huandu/go-sqlbuilder"
)

type WodSetRepository struct {
	db        *sql.DB
	dbTimeout time.Duration
}

func NewWodSetRepository(db *sql.DB, dbTimeout time.Duration) *WodSetRepository {
	return &WodSetRepository{
		db:        db,
		dbTimeout: dbTimeout,
	}
}

func (r *WodSetRepository) Save(ctx context.Context, wodSet crossfit.WodSet) error {
	wodSetSQLStruct := sqlbuilder.NewStruct(new(sqlWodSet))
	query, args := wodSetSQLStruct.InsertInto(sqlWodSetTable, sqlWodSet{
		ID:               wodSet.ID().String(),
		WodId:            wodSet.WodId().String(),
		SetNumber:        wodSet.SetNumber().Int(),
		BuyIn:            wodSet.SetBuyIn().Int(),
		BuyOut:           wodSet.SetBuyOut().Int(),
		EveryMinutes:     wodSet.SetEveryMinutes().Int(),
		RepsToAddByRound: wodSet.SetRepsToAddByRound().Int(),
		RestTime:         wodSet.SetRestTime().Int(),
		IsThen:           wodSet.SetIsThen().Bool(),
	}).Build()

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	_, err := r.db.ExecContext(ctxTimeout, query, args...)

	if err != nil {
		fmt.Printf("error saving wod_set: %v", err)
		return err
	}
	return nil
}
