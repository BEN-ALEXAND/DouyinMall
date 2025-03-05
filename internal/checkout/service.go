// internal/checkout/service.go
package checkout

import (
	"dymall/internal/payment"
	"dymall/internal/order"
)

type CheckoutService struct {
	orderService   *order.OrderService
	paymentService *payment.PaymentService
}

func NewCheckoutService(orderService *order.OrderService, paymentService *payment.PaymentService) *CheckoutService {
	return &CheckoutService{
		orderService:   orderService,
		paymentService: paymentService,
	}
}

func (s *CheckoutService) Checkout(userID uint32, items []order.CartItem, currency string, address order.Address, email string, creditCard payment.CreditCardInfo) (string, string, error) {
	// 创建订单
	order, err := s.orderService.PlaceOrder(userID, items)
	if err != nil {
		return "", "", err
	}

	// 支付
	transactionID, err := s.paymentService.Charge(order.ID, userID, creditCard, currency)
	if err != nil {
		return "", "", err
	}

	return order.ID, transactionID, nil
}