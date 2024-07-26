package mysql

import (
	"github.com/wodm8/wodm8-core/internal/members"
	"gorm.io/gorm"
)

type MembersRepository struct {
	db *gorm.DB
}

func NewMembersRepository(db *gorm.DB) *MembersRepository {
	return &MembersRepository{db: db}
}

func (r *MembersRepository) Save(member members.Members) error {
	result := r.db.Create(&member)
	return result.Error
}

func (r *MembersRepository) Find(id string) (members.Members, error) {
	var member members.Members
	result := r.db.Where("id = ?", id).First(&member)
	return member, result.Error
}

func (r *MembersRepository) FindByEmail(email string) (members.Members, error) {
	var member members.Members
	result := r.db.Where("email = ?", email).First(&member)
	if result.Error != nil {
		return member, result.Error
	}
	return member, nil
}

func (r *MembersRepository) Update(member members.Members) error {
	result := r.db.Save(&member)
	return result.Error
}
