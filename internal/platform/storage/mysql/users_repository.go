package mysql

import (
	"github.com/wodm8/wodm8-core/internal/users"
	"gorm.io/gorm"
)

type UsersRepository struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) *UsersRepository {
	return &UsersRepository{db: db}
}

func (r *UsersRepository) Save(user users.Users) error {
	result := r.db.Create(&user)
	return result.Error
}

func (r *UsersRepository) First(email string) (*users.Users, error) {
	var user users.Users
	result := r.db.First(&user, "email = ?", email)
	return &user, result.Error
}
