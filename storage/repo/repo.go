package repo

import (
	pb "github.com/muhriddinsalohiddin/online_store_order/genproto/order_service"
)

// OrderStorageI ...
type OrderStorageI interface {
	CreateOrder(pb.Order) (pb.Order, error)
	UpdateOrder(pb.Order) (pb.Order, error)
	GetOrderById(pb.GetOrderByIdReq) (pb.Order, error)
	DeleteById(pb.GetOrderByIdReq) error
	ListOrders(pb.ListOrderReq) (pb.ListOrderResp, error)
}
