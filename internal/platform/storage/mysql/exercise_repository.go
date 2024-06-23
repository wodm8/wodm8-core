package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/huandu/go-sqlbuilder"

	crossfit "github.com/wodm8/wodm8-core/internal"
)

// ExerciseRepository is a Postgresql integral.ExerciseRepository implementation
type ExerciseRepository struct {
	db *sql.DB
}

// NewExerciseRepository initializes a postgresql-based implementation of integral.ExerciseRepository
func NewExerciseRepository(db *sql.DB) *ExerciseRepository {
	return &ExerciseRepository{
		db: db,
	}
}

func (r *ExerciseRepository) Save(ctx context.Context, exercise crossfit.Exercise) error {
	fmt.Printf("ex_name: %s\n", exercise.Name())
	exerciseSQLStruct := sqlbuilder.NewStruct(new(sqlExercise))
	query, args := exerciseSQLStruct.InsertInto(sqlExerciseTable, sqlExercise{
		ID:   exercise.ID().String(),
		Name: exercise.Name().String(),
	}).Build()

	fmt.Printf("query: %s\n", query)
	fmt.Printf("args: %s\n", args)
	_, err := r.db.ExecContext(ctx, query, args...)
	fmt.Printf("err: %v\n", err)

	if err != nil {
		return fmt.Errorf("error saving exercise: %v", err)
	}
	return nil
}
