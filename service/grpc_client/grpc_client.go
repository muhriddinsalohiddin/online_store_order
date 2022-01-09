package grpc_client

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/muhriddinsalohiddin/online_store_order/config"
	pb "github.com/muhriddinsalohiddin/online_store_order/genproto/catalog_service"
)

// IGrpcClient ...
type IGrpcClient interface {
	CatalogService() pb.CatalogServiceClient
}

// GrpcClient ...
type GrpcClient struct {
	cfg            config.Config
	catalogService pb.CatalogServiceClient
}

// New ...
func New(cfg config.Config) (*GrpcClient, error) {
	connCatalog, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.CatalogServiceHost, cfg.CatalogServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("catalog service dial host: %s port: %d",
			cfg.CatalogServiceHost, cfg.CatalogServicePort)
	}
	grpcClient := &GrpcClient{
		cfg:            cfg,
		catalogService: pb.NewCatalogServiceClient(connCatalog),
	}
	return grpcClient, nil
}

// CatalogService ...
func (c *GrpcClient) CatalogService() pb.CatalogServiceClient {
	return c.catalogService
}
