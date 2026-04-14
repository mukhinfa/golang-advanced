package order

import (
	"gorm.io/gorm"

	"github.com/mukhinfa/golang-advanced/4-order-api/internal/product"
	"github.com/mukhinfa/golang-advanced/4-order-api/internal/user"
)

type Order struct {
	gorm.Model
	UserID   uint `json:"user_id"`
	User     user.User
	Products []product.Product `json:"products" gorm:"many2many:order_products;"`
}
