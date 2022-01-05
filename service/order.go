package service

import (
	"context"
	pb "github.com/muhriddinsalohiddin/online_store_order/genproto/order_service"
	"github.com/muhriddinsalohiddin/online_store_order/pkg/logger"
	"github.com/muhriddinsalohiddin/online_store_order/storage"
)

type OrderService struct {
	storage storage.IStorage
	logger logger.Logger
}

func NewOrderService(s storage.IStorage, l logger.Logger) *OrderService{
	return &OrderService{
		storage: s,
		logger: l,
	}
}

func (o OrderService) CreateOrder(context.Context, *pb.Order) (*pb.Order, error){
	return &pb.Order{},nil
}
func (o OrderService) UpdateOrder(context.Context, *pb.Order) (*pb.Order, error){
	return &pb.Order{},nil
}
func (o OrderService) GetOrderById(context.Context, *pb.GetOrderByIdReq) (*pb.Order, error){
	return &pb.Order{},nil
}
func (o OrderService) DeleteById(context.Context, *pb.GetOrderByIdReq) (*pb.EmptyResp, error){
	return &pb.EmptyResp{},nil
}
func (o OrderService) ListOrders(context.Context, *pb.ListOrderReq) (*pb.ListOrderResp, error){
	return &pb.ListOrderResp{},nil
}