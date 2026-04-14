package user

import (
	"errors"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type UserRepo interface {
	Create(user *User) (*User, error)
	GetByPhone(phone string) (*User, error)
	GetByID(id uint) (*User, error)
}

type userService struct {
	UserRepo
}

func NewService(userCreator UserRepo) *userService {
	return &userService{
		UserRepo: userCreator,
	}
}

func (s *userService) FindOrCreateByPhone(phone string) (*User, error) {
	user, err := s.UserRepo.GetByPhone(phone)
	if err == nil {
		return user, nil
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("user.FindOrCreateByPhone get: %w", err)
	}
	user, err = s.UserRepo.Create(&User{
		Phone: phone,
	})
	if err != nil {
		log.Printf("failed to create user: %v", err)
		return nil, err
	}
	return user, nil
}

func (s *userService) GetByID(id uint) (*User, error) {
	user, err := s.UserRepo.GetByID(id)
	if err != nil {
		log.Printf("failed to find user: %v", err)
		return nil, err
	}
	return user, nil
}
