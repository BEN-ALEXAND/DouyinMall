package product

type Product struct {
	ID          uint32
	Name        string
	Description string
	Price       float64
}

type ProductService struct {
	products []Product
}

func NewProductService() *ProductService {
	return &ProductService{
		products: []Product{
			{ID: 1, Name: "Product 1", Description: "Description 1", Price: 10.0},
			{ID: 2, Name: "Product 2", Description: "Description 2", Price: 20.0},
		},
	}
}

func (s *ProductService) GetProduct(id uint32) (*Product, error) {
	for _, p := range s.products {
		if p.ID == id {
			return &p, nil
		}
	}
	return nil, nil
}