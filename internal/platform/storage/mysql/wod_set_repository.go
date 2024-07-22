package mysql

import (
	"fmt"

	"github.com/wodm8/wodm8-core/internal/crossfit"
	"gorm.io/gorm"
)

type WodSetRepository struct {
	db *gorm.DB
}

func NewWodSetRepository(db *gorm.DB) *WodSetRepository {
	return &WodSetRepository{
		db: db,
	}
}

func (r *WodSetRepository) Save(wodSet crossfit.WodSet) error {
	result := r.db.Create(&wodSet)
	if result.Error != nil {
		fmt.Printf("error saving wod_set: %v", result.Error)
		return result.Error
	}
	return nil
}
