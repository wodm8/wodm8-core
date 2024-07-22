package mysql

import (
	"fmt"

	"github.com/wodm8/wodm8-core/internal/domain"
	"gorm.io/gorm"

	"github.com/wodm8/wodm8-core/internal/crossfit"
)

type WodRepository struct {
	db *gorm.DB
}

func NewWodRepository(db *gorm.DB) *WodRepository {
	return &WodRepository{
		db: db,
	}
}

func (r *WodRepository) Save(wod crossfit.Wod) error {
	result := r.db.Create(&wod)

	if result.Error != nil {
		fmt.Printf("error saving wod: %v", result.Error)
		return result.Error
	}
	return nil
}

func (r *WodRepository) Get(id string) ([]domain.CreatedWod, error) {
	var createdWods []domain.CreatedWod
	result := r.db.Where("wod_id = ?", id).Find(&createdWods)
	if result.Error != nil {
		fmt.Printf("error getting wod_data: %v", result.Error)
		return []domain.CreatedWod{}, result.Error
	}
	return createdWods, nil
}
