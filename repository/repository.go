package repository

import (
	"database/sql"
)

type Repository struct {
	User
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		User: NewUserPostgres(db),
	}
}
