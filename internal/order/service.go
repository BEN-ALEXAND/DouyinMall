package order

import (
	"time"
)

type Order struct {
	ID        string
	UserID    uint32
	Items     []CartItem
	Status    string
	CreatedAt int64
}

type OrderService struct {
	orders map[string]*Order
}

func NewOrderService() *OrderService {
	return &OrderService{
		orders: make(map[string]*Order),
	}
}

func (s *OrderService) PlaceOrder(userID uint32, items []CartItem) (*Order, error) {
	order := &Order{
		ID:        "order-123",
		UserID:    userID,
		Items:     items,
		Status:    "created",
		CreatedAt: time.Now().Unix(),
	}
	s.orders[order.ID] = order
	return order, nil
}