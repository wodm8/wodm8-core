package mysql

import (
	"fmt"

	"github.com/wodm8/wodm8-core/internal/box"
	"gorm.io/gorm"
)

type BoxRepository struct {
	db *gorm.DB
}

func NewBoxRepository(db *gorm.DB) *BoxRepository {
	return &BoxRepository{db: db}
}

func (r BoxRepository) Save(box box.Box) error {
	result := r.db.Create(&box)

	if result.Error != nil {
		fmt.Printf("error saving wod: %v", result.Error)
		return result.Error
	}
	return nil
}
