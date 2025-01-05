package repository

import (
	"database/sql"
	"ecommerce/models"
)

type Order interface {
	InsertOrder(
		order models.OrderCreateDTO,
		orderId string,
		orderItems []models.OrderItemCreateDTO,
	) error
}

type OrderPostgres struct {
	db *sql.DB
}

func NewOrderPostgres(db *sql.DB) *OrderPostgres {
	return &OrderPostgres{
		db,
	}
}
func (o *OrderPostgres) InsertOrder(
	order models.OrderCreateDTO,
	orderId string,
	orderItems []models.OrderItemCreateDTO,
) error {

	tx, err := o.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	_, err = tx.Exec(`
	INSERT INTO orders (id, user_id, total_price)
	VALUES ($1, $2, $3)
	`,
		orderId,
		order.UserId,
		order.TotalPrice,
	)

	if err != nil {
		return err
	}

	for _, orderItem := range orderItems {
		_, err = tx.Exec(`INSERT INTO order_items 
		(id, order_id, product_id, quantity, price)
		VALUES ($1, $2, $3, $4, $5)
		`,
			orderItem.Id,
			orderId,
			orderItem.ProductId,
			orderItem.Quantity,
			orderItem.Price,
		)
		if err != nil {
			return err
		}

		_, err = tx.Exec(`
		UPDATE products SET stock = stock - $1
		WHERE id = $2
		`,
			orderItem.Quantity,
			orderItem.ProductId,
		)
		if err != nil {
			return err
		}
	}
	_, err = tx.Exec(`DELETE FROM carts WHERE user_id = $1`, order.UserId)
	return err
}
