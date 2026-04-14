package order

import (
	"time"

	"github.com/mukhinfa/golang-advanced/4-order-api/internal/product"
)

type CreateOrderRequest struct {
	ProductIDs []uint `json:"product_ids" validate:"required,min=1"`
}

type OrderResponse struct {
	OrderID   uint              `json:"order_id"`
	Products  []product.Product `json:"products"`
	UserID    uint              `json:"user_id"`
	CreatedAt time.Time         `json:"created_at"`
}

type MyOrderResponse struct {
	Products  []product.Product `json:"products"`
	CreatedAt time.Time         `json:"created_at"`
}
