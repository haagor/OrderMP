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
		fmt.Println(err)
		return err
	}

	_, err = tx.Exec(`INSERT INTO orders(order_id, vat, total_price) VALUES($1, $2, $3)`, order.OrderID, order.VAT, order.TotalPrice)
	if err != nil {
		_ = tx.Rollback()
		fmt.Println(err)
		return err
	}

	for _, p := range order.Products {
		_, err := tx.Exec(`INSERT INTO products(product_id, name, price) VALUES($1, $2, $3)`, p.ProductID, p.Name, p.Price)
		if err != nil {
			_ = tx.Rollback()
			fmt.Println(err)
			return err
		}

		_, err = tx.Exec(`INSERT INTO order_to_product(order_id, product_id) VALUES($1, $2)`, order.OrderID, p.ProductID)
		if err != nil {
			_ = tx.Rollback()
			fmt.Println(err)
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
