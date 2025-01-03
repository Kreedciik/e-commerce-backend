package repository

import (
	"database/sql"
	"ecommerce/models"
	"time"

	"github.com/google/uuid"
)

type Product interface {
	InsertProduct(product models.CreateProductDTO) error
	UpdateProduct(product models.UpdateProductDTO) error
	DeleteProduct(id string) error
	GetProducts(name, description string) ([]models.Product, error)
}

type ProductPostgres struct {
	db *sql.DB
}

func NewProductPostgres(db *sql.DB) *ProductPostgres {
	return &ProductPostgres{
		db,
	}
}

func (p *ProductPostgres) InsertProduct(product models.CreateProductDTO) error {
	id := uuid.NewString()
	_, err := p.db.Exec(`INSERT INTO products 
	(id, name, description, price, stock)
	 VALUES ($1, $2, $3, $4, $5)`,
		id, product.Name, product.Description,
		product.Price, product.Stock,
	)
	return err
}

func (p *ProductPostgres) UpdateProduct(product models.UpdateProductDTO) error {
	_, err := p.db.Exec(`
	UPDATE products  SET 
	name = $1, 
	description = $2, 
	price = $3,
	stock = $4, 
	updated_at = $5
	WHERE id = $6`,
		product.Name, product.Description,
		product.Price, product.Stock,
		time.Now(), product.Id,
	)
	return err
}

func (p *ProductPostgres) DeleteProduct(id string) error {
	_, err := p.db.Exec(`DELETE FROM products WHERE id = $1`, id)
	return err
}

func (p *ProductPostgres) GetProducts(name, description string) ([]models.Product, error) {
	products := []models.Product{}
	rows, err := p.db.Query(`SELECT id, name, description, price, stock
	FROM products
	WHERE name ILIKE $1 OR description ILIKE $2
	ORDER BY created_at
	`, "%"+name+"%", "%"+description+"%")

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
