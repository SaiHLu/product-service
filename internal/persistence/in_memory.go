package persistence

import (
	"SaiHLu/proto/protogen/go/product"
)

type InMemoryStore struct {
	products map[int]*product.Product
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		products: make(map[int]*product.Product),
	}
}

func (s *InMemoryStore) CreateProduct(req *product.CreateProductRequest) (*product.CreateProductResponse, error) {
	id := len(s.products) + 1
	s.products[id] = req.Product

	return &product.CreateProductResponse{
		Product: s.products[id],
	}, nil
}

func (s *InMemoryStore) GetProductList(req *product.GetProductListRequest) (*product.GetProductListResponse, error) {
	products := make([]*product.Product, 0, len(s.products))
	for _, p := range s.products {
		products = append(products, p)
	}

	return &product.GetProductListResponse{
		Products: products,
	}, nil
}
