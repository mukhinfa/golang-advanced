package order

import (
	"fmt"

	"github.com/mukhinfa/golang-advanced/4-order-api/pkg/db"
)

type orderRepository struct {
	database *db.Db
}

func NewRepository(db *db.Db) *orderRepository {
	return &orderRepository{
		database: db,
	}
}

func (r *orderRepository) Create(order *Order) error {
	result := r.database.Create(order)
	if result.Error != nil {
		return fmt.Errorf("order.Repository.Create failed to create order: %w", result.Error)
	}
	return nil
}
func (r *orderRepository) GetByOrderID(id uint) (*Order, error) {
	var order Order
	result := r.database.
		Preload("Products").
		Preload("User").
		First(&order, id)

	if result.Error != nil {
		return nil, fmt.Errorf("order.Repository.GetByID failed to get order by ID: %w", result.Error)
	}
	return &order, nil
}
func (r *orderRepository) GetByUserID(userID uint, limit, offset int) ([]Order, error) {
	var orders []Order
	result := r.database.
		Preload("Products").
		Where("user_id = ?", userID).
		Order("created_at asc").
		Limit(limit).
		Offset(offset).
		Find(&orders)

	if result.Error != nil {
		return nil, fmt.Errorf("order.Repository.GetByUserID: %w", result.Error)
	}
	return orders, nil
}
