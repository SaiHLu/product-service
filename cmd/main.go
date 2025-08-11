package main

import (
	"log"

	"github.com/SaiHLu/product-service/internal/app"
	grpcserver "github.com/SaiHLu/product-service/internal/transport/grpc"
)

func main() {
	appConf := app.NewAppConfig()
	appConf.Setup()

	if err := grpcserver.Run(appConf); err != nil {
		log.Fatalf("error: %+v", err)
	}
}
