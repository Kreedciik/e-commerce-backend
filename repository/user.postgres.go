package repository

import (
	"database/sql"
	"ecommerce/models"

	"github.com/google/uuid"
)

// All methods of UserPostgres should be define here
type User interface {
	InsertUser(user models.CreateUserDTO) error
	UpdateUser(user models.UpdateUserDTO) error
	DeleteUser(id string) error
	FindUserByID(id string) (models.User, error)
}

type UserPostgres struct {
	db *sql.DB
}

func NewUserPostgres(db *sql.DB) *UserPostgres {
	return &UserPostgres{
		db,
	}
}

func (u *UserPostgres) InsertUser(user models.CreateUserDTO) error {
	id := uuid.NewString()
	_, err := u.db.Exec(`INSERT INTO users (id, name, role, email)
	 VALUES ($1, $2, $3, $4)`,
		id, user.Name, user.Role, user.Email)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserPostgres) UpdateUser(user models.UpdateUserDTO) error {
	_, err := u.db.Exec(`
	UPDATE users  SET 
	name = $1, role = $2, email = $3
	WHERE id = $1`,
		user.Name, user.Role,
		user.Email, user.Id)
	return err
}

func (u *UserPostgres) DeleteUser(id string) error {
	// Implement delete in database
	return nil
}

func (u *UserPostgres) FindUserByID(id string) (models.User, error) {
	// Implement retriving user by ID
	return models.User{}, nil
}
