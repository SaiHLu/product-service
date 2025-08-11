package app

import (
	"github.com/SaiHLu/product-service/internal/persistence"
	"github.com/SaiHLu/product-service/internal/service"
)

type AppConfig struct {
	Store          persistence.Store
	ProductService *service.ProductService
}

var singletonAppConfig *AppConfig

func NewAppConfig() *AppConfig {
	if singletonAppConfig != nil {
		return singletonAppConfig
	}

	return &AppConfig{}
}
