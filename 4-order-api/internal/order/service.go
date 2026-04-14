package order

import (
	"github.com/mukhinfa/golang-advanced/4-order-api/internal/product"
	"github.com/mukhinfa/golang-advanced/4-order-api/internal/user"
)

type OrderRepository interface {
	Create(o *Order) error
	GetByOrderID(id uint) (*Order, error)
	GetByUserID(userID uint, limit, offset int) ([]Order, error)
}

type ProductProvider interface {
	GetByID(id uint) (*product.Product, error)
}

type UserProvider interface {
	FindOrCreateByPhone(phone string) (*user.User, error)
}

type orderService struct {
	orders   OrderRepository
	products ProductProvider
	users    UserProvider
}

func NewService(
	orders OrderRepository,
	products ProductProvider,
	users UserProvider,
) *orderService {
	return &orderService{orders: orders, products: products, users: users}
}
func (s *orderService) CreateOrder(phone string, req CreateOrderRequest) (*Order, error) {
	user, err := s.users.FindOrCreateByPhone(phone)
	if err != nil {
		return nil, err
	}
	products := make([]product.Product, 0, len(req.ProductIDs))
	for _, v := range req.ProductIDs {
		product, err := s.products.GetByID(v)
		if err != nil {
			return nil, err
		}
		products = append(products, *product)
	}
	order := &Order{
		UserID:   user.ID,
		Products: products,
	}
	err = s.orders.Create(order)
	if err != nil {
		return nil, err
	}
	return order, nil
}
func (s *orderService) GetOrder(id uint) (*Order, error) {
	order, err := s.orders.GetByOrderID(id)
	if err != nil {
		return nil, err
	}
	return order, nil
}
func (s *orderService) GetUserOrders(phone string, limit, offset int) ([]Order, error) {
	user, err := s.users.FindOrCreateByPhone(phone)
	if err != nil {
		return nil, err
	}
	orders, err := s.orders.GetByUserID(user.ID, limit, offset)
	if err != nil {
		return nil, err
	}
	return orders, nil
}
