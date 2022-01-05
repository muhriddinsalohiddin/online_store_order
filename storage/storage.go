package storage

import (
	"github.com/jmoiron/sqlx"
	"github.com/muhriddinsalohiddin/online_store_order/storage/repo"
	"github.com/muhriddinsalohiddin/online_store_order/storage/postgres"

)

// IStorage ...
type IStorage interface {
	Order() repo.OrderStorageI
}

type storagePg struct {
	db *sqlx.DB
	orderRepo repo.OrderStorageI
}

func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{
		db: db,
		orderRepo: postgres.NewOrderRepo(db),
	}
}
func (s storagePg) Order() repo.OrderStorageI{
	return s.orderRepo
}