package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	// Name  string `json:"name" gorm:"not null"`
	Phone string `json:"phone" gorm:"uniqueIndex;not null"`
}
