package crossfit

import (
	"github.com/wodm8/wodm8-core/internal/domain"
)

type Wod struct {
	ID        string `gorm:"primaryKey"`
	Name      string
	WodDate   string
	WodTypeId uint8
}

func NewWod(id, wodName, date string, wodTypeId uint8) (Wod, error) {

	return Wod{
		ID:        id,
		Name:      wodName,
		WodDate:   date,
		WodTypeId: wodTypeId,
	}, nil
}

type WodRepository interface {
	Save(wod Wod) error
	Get(id string) ([]domain.CreatedWod, error)
}

//go:generate mockery --case=snake --outpkg=storagemocks --output=../platform/storage/storagemocks --name=WodRepository
