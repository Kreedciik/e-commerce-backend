package repository

import (
	"database/sql"
)

type Repository struct {
	User
	Product
	Cart
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		User:    NewUserPostgres(db),
		Product: NewProductPostgres(db),
		Cart:    NewCartPostgres(db),
	}
}
