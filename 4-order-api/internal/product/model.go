package product

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string         `json:"name" gorm:"not null"`
	Description string         `json:"description,omitempty"`
	Images      pq.StringArray `json:"images,omitempty" gorm:"type:jsonb"`
}
