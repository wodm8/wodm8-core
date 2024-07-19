package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/wodm8/wodm8-core/internal/crossfit"

	"github.com/huandu/go-sqlbuilder"
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
		ID:              exerciseWod.ID().String(),
		WodID:           exerciseWod.WodID().String(),
		ExerciseID:      exerciseWod.ExerciseID().Int(),
		RoundNumber:     exerciseWod.RoundNumber().Int(),
		SetNumber:       exerciseWod.SetNumber().Int(),
		Repetitions:     exerciseWod.Repetitions().Int(),
		RepetitionsUnit: exerciseWod.RepetitionsUnit().String(),
		Weight:          exerciseWod.Weight().Float(),
		WeightUnit:      exerciseWod.WeightUnit().String(),
		Distance:        exerciseWod.Distance().Float(),
		DistanceUnit:    exerciseWod.DistanceUnit().String(),
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
