package persistence

import "SaiHLu/proto/protogen/go/product"

type Store interface {
	CreateProduct(*product.CreateProductRequest) (*product.CreateProductResponse, error)
	GetProductList(*product.GetProductListRequest) (*product.GetProductListResponse, error)
}
