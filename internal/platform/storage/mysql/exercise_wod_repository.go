package mysql

import (
	"fmt"

	"github.com/wodm8/wodm8-core/internal/crossfit"
	"gorm.io/gorm"
)

type ExerciseWodRepository struct {
	db *gorm.DB
}

func NewExerciseWodRepository(db *gorm.DB) *ExerciseWodRepository {
	return &ExerciseWodRepository{
		db: db,
	}
}

func (r *ExerciseWodRepository) Save(exerciseWod crossfit.ExerciseWod) error {
	result := r.db.Create(&exerciseWod)

	if result.Error != nil {
		fmt.Printf("error saving exercise wod: %v", result.Error)
		return result.Error
	}
	return nil
}
