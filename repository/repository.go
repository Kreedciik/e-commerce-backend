package repository

import (
	"database/sql"
)

type Repository struct {
	User
	Product
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		User:    NewUserPostgres(db),
		Product: NewProductPostgres(db),
	}
}
