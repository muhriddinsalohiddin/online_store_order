package service

import (
	"context"

	"github.com/gofrs/uuid"
	pb "github.com/muhriddinsalohiddin/online_store_order/genproto/order_service"
	"github.com/muhriddinsalohiddin/online_store_order/pkg/logger"
	"github.com/muhriddinsalohiddin/online_store_order/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type OrderService struct {
	storage storage.IStorage
	logger  logger.Logger
}

func NewOrderService(s storage.IStorage, l logger.Logger) *OrderService {
	return &OrderService{
		storage: s,
		logger:  l,
	}
}

func (o OrderService) CreateOrder(ctx context.Context, in *pb.Order) (*pb.Order, error) {
	id, err := uuid.NewV4()
	if err != nil {
		o.logger.Error("failed to generate uuid", logger.Error(err))
		return nil, status.Error(codes.Internal, "failed to generate uuid")
	}
	in.Id = id.String()
	order, err := o.storage.Order().CreateOrder(*in)
	if err != nil {
		o.logger.Error("failed to create order",logger.Error(err))
		return nil, status.Error(codes.Internal, "failed to create order")
	}
	return &order, nil
}

func (o OrderService) UpdateOrder(ctx context.Context, in *pb.Order) (*pb.Order, error) {
	order,err := o.storage.Order().UpdateOrder(*in)
	if err != nil {
		o.logger.Error("failed to update order",logger.Error(err))
		return nil, status.Error(codes.Internal, "failed to update order")
	}
	return &order, nil
}

func (o OrderService) GetOrderById(ctx context.Context, in *pb.GetOrderByIdReq) (*pb.Order, error) {
	order,err := o.storage.Order().GetOrderById(*in)
	if err != nil {
		o.logger.Error("failed to get order",logger.Error(err))
		return nil, status.Error(codes.Internal, "failed to get order")
	}
	return &order, nil
}

func (o OrderService) DeleteById(ctx context.Context, in *pb.GetOrderByIdReq) (*pb.EmptyResp, error) {
	err := o.storage.Order().DeleteById(*in)
	if err != nil {
		o.logger.Error("failed to delete order",logger.Error(err))
		return nil, status.Error(codes.Internal, "failed to delete order")
	}
	return &pb.EmptyResp{}, nil
}

func (o OrderService) ListOrders(ctx context.Context, in *pb.ListOrderReq) (*pb.ListOrderResp, error) {
	order,err := o.storage.Order().ListOrders(*in)
	if err != nil {
		o.logger.Error("failed to get list orders",logger.Error(err))
		return nil, status.Error(codes.Internal, "failed to get list orders")
	}
	return &order, nil
}
