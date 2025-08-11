package app

import (
	"github.com/SaiHLu/product-service/internal/persistence"
	"github.com/SaiHLu/product-service/internal/service"
)

func (conf *AppConfig) Setup() {
	conf.Store = persistence.NewInMemoryStore()
	conf.ProductService = service.NewProductService(conf.Store)
}
