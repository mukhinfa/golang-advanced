package user

import (
	"fmt"

	"github.com/mukhinfa/golang-advanced/4-order-api/pkg/db"
)

type userRepository struct {
	database *db.Db
}

func NewRepository(database *db.Db) *userRepository {
	return &userRepository{
		database: database,
	}
}

func (r *userRepository) Create(user *User) (*User, error) {
	result := r.database.Create(user)
	if result.Error != nil {
		return nil, fmt.Errorf("user.Repository.Create failed to create user: %w", result.Error)
	}
	return user, nil
}
func (r *userRepository) GetByPhone(phone string) (*User, error) {
	var user User
	result := r.database.First(&user, "phone = ?", phone)
	if result.Error != nil {
		return nil, fmt.Errorf("user.Repository.GetByPhone failed to get user by phone: %w", result.Error)
	}
	return &user, nil
}
func (r *userRepository) GetByID(id uint) (*User, error) {
	var user User
	result := r.database.First(&user, id)
	if result.Error != nil {
		return nil, fmt.Errorf("user.Repository.GetByID failed to get user by ID: %w", result.Error)
	}
	return &user, nil
}
