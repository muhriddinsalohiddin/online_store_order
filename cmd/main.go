package main

import (
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/muhriddinsalohiddin/online_store_order/config"
	pb "github.com/muhriddinsalohiddin/online_store_order/genproto/order_service"
	"github.com/muhriddinsalohiddin/online_store_order/pkg/db"
	"github.com/muhriddinsalohiddin/online_store_order/pkg/logger"
	"github.com/muhriddinsalohiddin/online_store_order/service"
	grpcClient "github.com/muhriddinsalohiddin/online_store_order/service/grpc_client"
	"github.com/muhriddinsalohiddin/online_store_order/storage"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.LogLevel, "order_service")

	defer func(l logger.Logger) {
		err := logger.Cleanup(l)
		if err != nil {
			log.Fatal("failed cleanup logger", logger.Error(err))
		}
	}(log)

	log.Info("main: sqlxConfig",
		logger.String("host", cfg.PostgresHost),
		logger.Int("port", cfg.PostgresPort),
		logger.String("database", cfg.PostgresDatabase),
	)

	connDB, err := db.ConnectToDB(cfg)
	if err != nil {
		log.Fatal("failed to connect db", logger.Error(err))
	}
	pgStorage := storage.NewStoragePg(connDB)

	grpcClient, err := grpcClient.New(cfg)
	if err != nil {
		log.Error("error establishing grpc connection", logger.Error(err))
		return
	}

	orderService := service.NewOrderService(pgStorage, log, grpcClient)

	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Fatal("Connection GRPC error", logger.Error(err))
	}

	s := grpc.NewServer()

	pb.RegisterOrderServiceServer(s, orderService)
	reflection.Register(s)
	log.Info("Main: server running",
		logger.String("port", cfg.RPCPort))
	if err := s.Serve(lis); err != nil {
		log.Fatal("Error while listening", logger.Error(err))
	}
}
