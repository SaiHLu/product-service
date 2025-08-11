package grpcserver

import (
	pb "SaiHLu/proto/protogen/go/product"
	"context"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/SaiHLu/product-service/internal/app"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

func Run(appConfig *app.AppConfig) error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	grpcServer := grpc.NewServer()

	// inMemoStore := persistence.NewInMemoryStore()
	// productService := service.NewProductService(inMemoStore)
	pb.RegisterProductServiceServer(grpcServer, appConfig.ProductService)

	errg, ctx := errgroup.WithContext(ctx)

	errg.Go(func() error {
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			return err
		}

		log.Println("Product Service Server is running on port: 50051")
		return grpcServer.Serve(lis)
	})

	errg.Go(func() error {
		<-ctx.Done()

		grpcServer.GracefulStop()
		return nil
	})

	return errg.Wait()
}
