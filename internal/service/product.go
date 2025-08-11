package service

import (
	"SaiHLu/proto/protogen/go/product"
	"context"

	"github.com/SaiHLu/product-service/internal/persistence"
)

type ProductService struct {
	product.UnimplementedProductServiceServer

	store persistence.Store
}

func NewProductService(store persistence.Store) *ProductService {
	return &ProductService{
		store: store,
	}
}

func (p *ProductService) CreateProduct(ctx context.Context, req *product.CreateProductRequest) (*product.CreateProductResponse, error) {
	return p.store.CreateProduct(req)
}

func (p *ProductService) GetProductList(ctx context.Context, req *product.GetProductListRequest) (*product.GetProductListResponse, error) {
	return p.store.GetProductList(req)
}
