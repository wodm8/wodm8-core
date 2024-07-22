package mysql

import (
	"fmt"

	"github.com/wodm8/wodm8-core/internal/crossfit"
	"gorm.io/gorm"
)

type WodRoundRepository struct {
	db *gorm.DB
}

func NewWodRoundRepository(db *gorm.DB) *WodRoundRepository {
	return &WodRoundRepository{
		db: db,
	}
}

func (r *WodRoundRepository) Save(wodRound crossfit.WodRound) error {
	result := r.db.Create(&wodRound)
	if result.Error != nil {
		fmt.Printf("error saving wod round: %v", result.Error)
		return result.Error
	}
	return nil
}
