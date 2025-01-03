package repository

import (
	"database/sql"
	"ecommerce/models"

	"github.com/google/uuid"
)

type Cart interface {
	InsertToCart(product models.PutToCartDTO) error
	RemoveFromCart(userId, productId string) error
	FindAllProductsFromCart(userId string) ([]models.Product, error)
}

type CartPostgres struct {
	db *sql.DB
}

func NewCartPostgres(db *sql.DB) *CartPostgres {
	return &CartPostgres{
		db,
	}
}

func (c *CartPostgres) InsertToCart(cart models.PutToCartDTO) error {

	id := uuid.NewString()
	_, err := c.db.Exec(`INSERT INTO carts 
	(id, user_id, product_id, quantity)
	 VALUES ($1, $2, $3, $4)`,
		id,
		cart.UserId,
		cart.ProductId,
		cart.Quantity,
	)
	return err
}

func (c *CartPostgres) RemoveFromCart(userId, productId string) error {
	_, err := c.db.Exec(`
	DELETE FROM carts WHERE user_id = $1 AND product_id = $2`,
		userId,
		productId,
	)
	return err
}

func (c *CartPostgres) FindAllProductsFromCart(userId string) ([]models.Product, error) {
	products := []models.Product{}
	rows, err := c.db.Query(`
	SELECT p.id, p.name, p.description, p.price, p.stock
	FROM products p
	INNER JOIN carts c ON c.product_id = p.product_id
	WHERE c.user_id = $1
	`, userId)

	if err != nil {
		return products, err
	}
	defer rows.Close()

	for rows.Next() {
		product := models.Product{}
		err := rows.Scan(&product.Id, &product.Name,
			&product.Description, &product.Price,
			&product.Stock,
		)

		if err != nil {
			return products, err
		}

		products = append(products, product)

	}

	return products, nil
}
