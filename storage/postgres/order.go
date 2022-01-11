package postgres

import (
	"database/sql"
	"time"

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

func (o *orderRepo) CreateOrder(in pb.Order) (pb.Order, error) {
	err := o.db.QueryRow(`
		INSERT INTO orders (id,book_id,description,created_at,updated_at)
		VALUES ($1,$2,$3,$4,$5)`,
		in.Id,
		in.BookId,
		in.Description,
		time.Now().UTC(),
		time.Now().UTC(),
	).Err()
	if err != nil {
		return pb.Order{}, err
	}

	order, err := o.GetOrderById(pb.GetOrderByIdReq{Id: in.Id})
	if err != nil {
		return pb.Order{}, err
	}

	return order, nil
}

func (o *orderRepo) UpdateOrder(in pb.Order) (pb.Order, error) {
	result, err := o.db.Exec(`
		UPDATE orders
		SET
			description = $1,
			updated_at = $2
		WHERE id = $3`,
		in.Description,
		time.Now().UTC(),
		in.Id,
	)
	if err != nil {
		return pb.Order{}, err
	}
	if i, _ := result.RowsAffected(); i == 0 {
		return pb.Order{}, sql.ErrNoRows
	}

	order, err := o.GetOrderById(pb.GetOrderByIdReq{Id: in.Id})
	if err != nil {
		return pb.Order{}, err
	}
	return order, nil
}

func (o *orderRepo) GetOrderById(in pb.GetOrderByIdReq) (pb.Order, error) {
	var order pb.Order
	err := o.db.QueryRow(`
		SELECT
			id,
			book_id,
			description,
			created_at,
			updated_at
		FROM orders
		WHERE id = $1
		AND deleated_at IS NULL`, in.Id,
	).Scan(
		&order.Id,
		&order.BookId,
		&order.Description,
		&order.CreatedAt,
		&order.UpdatedAt,
	)
	if err != nil {
		return pb.Order{}, err
	}

	return order, nil
}

func (o *orderRepo) DeleteById(in pb.GetOrderByIdReq) error {
	result, err := o.db.Exec(`
		UPDATE orders
		SET deleated_at = $1
		WHERE id = $2`,
		time.Now().UTC(),
		in.Id,
	)
	if err != nil {
		return err
	}
	if i, _ := result.RowsAffected(); i == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (o *orderRepo) ListOrders(in pb.ListOrderReq) (pb.ListOrderResp, error) {
	offset := (in.Page - 1) * in.Limit
	rows, err := o.db.Query(`
		SELECT
			id,
			book_id,
			description,
			created_at,
			updated_at
		FROM orders
		WHERE deleated_at IS NULL
		LIMIT $1
		OFFSET $2`,
		in.Limit,
		offset,
	)
	if err != nil {
		return pb.ListOrderResp{}, err
	}
	defer rows.Close()
	var orders pb.ListOrderResp
	for rows.Next() {
		var order pb.Order
		err = rows.Scan(
			&order.Id,
			&order.BookId,
			&order.Description,
			&order.CreatedAt,
			&order.UpdatedAt,
		)
		if err != nil {
			return pb.ListOrderResp{}, err
		}
		orders.Orders = append(orders.Orders, &order)
	}
	err = o.db.QueryRow(`
		SELECT COUNT(*)
		FROM orders
		WHERE deleated_at IS NULL
	`).Scan(&orders.Count)
	if err != nil {
		return pb.ListOrderResp{}, err
	}
	return orders, nil
}
