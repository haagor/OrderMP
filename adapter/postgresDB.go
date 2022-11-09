package adapter

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/haagor/orderMP/model"
)

const (
	Host     = "localhost"
	Port     = 5432
	User     = "goadapter"
	Password = "goadapter"
	Dbname   = "orderdb"
)

type PostgresAdapter struct {
	DB *sql.DB
}

func (pa PostgresAdapter) AddOrderWithProduct(order model.Order) error {
	ctx := context.Background()
	tx, err := pa.DB.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("AddOrderWithProduct: error setup connection to psql: %w", err)
	}

	_, err = tx.Exec(`INSERT INTO orders(order_id, vat, total_price) VALUES($1, $2, $3)`, order.OrderID, order.VAT, order.TotalPrice)
	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("AddOrderWithProduct: error insert order: %w", err)
	}

	for _, p := range order.Products {
		_, err := tx.Exec(`INSERT INTO products(product_id, name, price) VALUES($1, $2, $3) ON CONFLICT DO NOTHING`, p.ProductID, p.Name, p.Price)
		if err != nil {
			_ = tx.Rollback()
			return fmt.Errorf("AddOrderWithProduct: error insert product: %w", err)
		}

		_, err = tx.Exec(`INSERT INTO order_to_product(order_id, product_id) VALUES($1, $2)`, order.OrderID, p.ProductID)
		if err != nil {
			_ = tx.Rollback()
			return fmt.Errorf("AddOrderWithProduct: error insert order_to_product: %w", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("AddOrderWithProduct: error commit query: %w", err)
	}

	return nil
}
