package main

import (
	"net"

	"github.com/muhriddinsalohiddin/online_store_order/config"
	pb "github.com/muhriddinsalohiddin/online_store_order/genproto/order_service"
	"github.com/muhriddinsalohiddin/online_store_order/pkg/db"
	"github.com/muhriddinsalohiddin/online_store_order/pkg/logger"
	"github.com/muhriddinsalohiddin/online_store_order/service"
	"github.com/muhriddinsalohiddin/online_store_order/storage"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
		logger.String("host",cfg.PostgresHost),
		logger.Int("port",cfg.PostgresPort),
		logger.String("database",cfg.PostgresDatabase),
	)

	connDB, err := db.ConnectToDB(cfg)

	if err != nil {
		log.Fatal("failed to connect db", logger.Error(err))
	}
	pgStorage := storage.NewStoragePg(connDB)

	orderService := service.NewOrderService(pgStorage,log)

	lis,err := net.Listen("tcp",cfg.RPCPort)
	if err != nil {
		log.Fatal("Connection GRPC error",logger.Error(err))
	}

	s := grpc.NewServer()

	pb.RegisterOrderServiceServer(s, orderService)
	reflection.Register(s)
	log.Info("Main: server running",
		logger.String("port",cfg.RPCPort))	
	if err := s.Serve(lis);err != nil {
		log.Fatal("Error while listening",logger.Error(err))
	}
}
