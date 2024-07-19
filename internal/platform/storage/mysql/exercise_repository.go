package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/wodm8/wodm8-core/internal/crossfit"

	"github.com/huandu/go-sqlbuilder"
)

// ExerciseRepository is a Postgresql integral.ExerciseRepository implementation
type ExerciseRepository struct {
	db        *sql.DB
	dbTimeout time.Duration
}

// NewExerciseRepository initializes a postgresql-based implementation of integral.ExerciseRepository
func NewExerciseRepository(db *sql.DB, dbTimeout time.Duration) *ExerciseRepository {
	return &ExerciseRepository{
		db:        db,
		dbTimeout: dbTimeout,
	}
}

func (r *ExerciseRepository) Save(ctx context.Context, exercise crossfit.Exercise) error {
	exerciseSQLStruct := sqlbuilder.NewStruct(new(sqlExercise))
	query, args := exerciseSQLStruct.InsertInto(sqlExerciseTable, sqlExercise{
		Name: exercise.Name().String(),
	}).Build()

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	_, err := r.db.ExecContext(ctxTimeout, query, args...)

	if err != nil {
		return fmt.Errorf("error saving exercise: %v", err)
	}
	return nil
}
