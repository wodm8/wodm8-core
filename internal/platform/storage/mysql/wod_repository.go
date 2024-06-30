package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/huandu/go-sqlbuilder"
	crossfit "github.com/wodm8/wodm8-core/internal"
)

type WodRepository struct {
	db        *sql.DB
	dbTimeout time.Duration
}

func NewWodRepository(db *sql.DB, dbTimeout time.Duration) *WodRepository {
	return &WodRepository{
		db:        db,
		dbTimeout: dbTimeout,
	}
}

func (r *WodRepository) Save(ctx context.Context, wod crossfit.Wod) error {
	wodSQLStruct := sqlbuilder.NewStruct(new(sqlWod))
	query, args := wodSQLStruct.InsertInto(sqlWodTable, sqlWod{
		ID:             wod.ID().String(),
		Name:           wod.Name().String(),
		Rounds:         wod.Rounds().Int(),
		NumberSections: wod.NumberSections().Int(),
		TimerTypeId:    wod.TimerType().Int(),
	}).Build()

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	_, err := r.db.ExecContext(ctxTimeout, query, args...)

	if err != nil {
		fmt.Printf("error saving wod: %v", err)
		return err
	}
	return nil
}
