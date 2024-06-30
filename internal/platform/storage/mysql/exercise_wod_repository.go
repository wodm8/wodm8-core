package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/huandu/go-sqlbuilder"
	crossfit "github.com/wodm8/wodm8-core/internal"
)

type ExerciseWodRepository struct {
	db        *sql.DB
	dbTimeout time.Duration
}

func NewExerciseWodRepository(db *sql.DB, dbTimeout time.Duration) *ExerciseWodRepository {
	return &ExerciseWodRepository{
		db:        db,
		dbTimeout: dbTimeout,
	}
}

func (r *ExerciseWodRepository) Save(ctx context.Context, exerciseWod crossfit.ExerciseWod) error {
	exerciseWodSQLStruct := sqlbuilder.NewStruct(new(sqlExerciseWod))
	query, args := exerciseWodSQLStruct.InsertInto(sqlExerciseWodTable, sqlExerciseWod{
		ID:                 exerciseWod.ID().String(),
		WodID:              exerciseWod.WodID().String(),
		ExerciseID:         exerciseWod.ExerciseID().String(),
		Reps:               exerciseWod.Repetitions().Int(),
		Weight:             exerciseWod.Weight().Float(),
		WeightUnit:         exerciseWod.WeightUnit().String(),
		WodSection:         exerciseWod.Section().Int(),
		SectionTimerTypeID: exerciseWod.TimerType().Int(),
		SectionCap:         exerciseWod.Cap().Int(),
	}).Build()

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	_, err := r.db.ExecContext(ctxTimeout, query, args...)

	if err != nil {
		fmt.Printf("error saving exercise wod: %v", err)
		return err
	}
	return nil
}
