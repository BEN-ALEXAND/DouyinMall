package payment

type CreditCardInfo struct {
	Number          string
	ExpirationYear  int32
	ExpirationMonth int32
	CVV             int32
}

type PaymentService struct{}

func NewPaymentService() *PaymentService {
	return &PaymentService{}
}

func (s *PaymentService) Charge(orderID string, userID uint32, creditCard CreditCardInfo, currency string) (string, error) {
	return "txn-123456", nil
}
