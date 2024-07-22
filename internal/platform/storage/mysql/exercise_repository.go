package mysql

import (
	"fmt"

	"github.com/wodm8/wodm8-core/internal/crossfit"
	"gorm.io/gorm"
)

type ExerciseRepository struct {
	db *gorm.DB
}

func NewExerciseRepository(db *gorm.DB) *ExerciseRepository {
	return &ExerciseRepository{
		db: db,
	}
}

func (r *ExerciseRepository) Save(exercise crossfit.Exercise) error {
	result := r.db.Create(&exercise)
	if result.Error != nil {
		return fmt.Errorf("error saving exercise: %v", result.Error)
	}

	return nil
}
