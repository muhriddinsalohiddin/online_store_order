package postgres

import (
	"github.com/jmoiron/sqlx"
	pb "github.com/muhriddinsalohiddin/online_store_order/genproto/order_service"
)

type orderRepo struct {
	db *sqlx.DB
}

// NewOrderRepo
func NewOrderRepo(db *sqlx.DB) *orderRepo {
	return &orderRepo{db: db}
}

func (o *orderRepo) CreateOrder(pb.Order) (pb.Order, error){
	return pb.Order{},nil
}
func (o *orderRepo) UpdateOrder(pb.Order) (pb.Order, error){
	return pb.Order{},nil
}
func (o *orderRepo) GetOrderById(pb.GetOrderByIdReq) (pb.Order, error){
	return pb.Order{},nil
}
func (o *orderRepo) DeleteById(pb.GetOrderByIdReq) (pb.EmptyResp, error){
	return pb.EmptyResp{},nil
}
func (o *orderRepo) ListOrders(pb.ListOrderReq) (pb.ListOrderResp, error){
	return pb.ListOrderResp{},nil
}