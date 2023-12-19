package database

import (
	"database/sql"

	"github.com/malandrim/goexpert/20-CleanArchitecture/internal/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{Db: db}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	stmt, err := r.Db.Prepare("INSERT INTO orders (id, price, tax, final_price) VALUES (? , ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(order.id, order.price, order.tax, order.final_price)
	if err != nil {
		return err
	}
	return nil
}
