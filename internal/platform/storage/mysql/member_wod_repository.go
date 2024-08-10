package mysql

import (
	"github.com/wodm8/wodm8-core/internal/crossfit"
	"gorm.io/gorm"
)

type MemberWodRepository struct {
	db *gorm.DB
}

func NewMemberWodRepository(db *gorm.DB) *MemberWodRepository {
	return &MemberWodRepository{db: db}
}

func (r *MemberWodRepository) Save(memberWod crossfit.MemberWod) error {
	result := r.db.Create(&memberWod)
	return result.Error
}
