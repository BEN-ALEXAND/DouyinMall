package cart

type CartItem struct {
	ProductID uint32
	Quantity  int32
}

type Cart struct {
	UserID uint32
	Items  []CartItem
}

type CartService struct {
	carts map[uint32]*Cart
}

func NewCartService() *CartService {
	return &CartService{
		carts: make(map[uint32]*Cart),
	}
}

func (s *CartService) AddItem(userID uint32, item CartItem) {
	if cart, exists := s.carts[userID]; exists {
		cart.Items = append(cart.Items, item)
	} else {
		s.carts[userID] = &Cart{UserID: userID, Items: []CartItem{item}}
	}
}