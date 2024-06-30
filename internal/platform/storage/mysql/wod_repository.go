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

	fmt.Printf("**query: %v\n", query)
	fmt.Printf("**args: %v\n", args)
	result, err := r.db.ExecContext(ctxTimeout, `-- name: InsertWod :execresult
INSERT INTO WOD (id, wod_name, rounds, number_sections, timer_type_id) VALUES (?, ?, ?, ?, ?)`, wod.ID().String(), wod.Name().String(), wod.Rounds().Int(), wod.NumberSections().Int(), wod.TimerType().Int())

	fmt.Printf("result: %v\n", result)
	fmt.Printf("error saving wod: %v\n", err)
	if err != nil {
		return fmt.Errorf("error saving wod: %v", err)
	}
	return nil
}
