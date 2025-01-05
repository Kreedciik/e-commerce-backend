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
	FindUserByEmail(email string) (models.User, error)
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
	_, err := u.db.Exec(`INSERT INTO users (id, name, role, email, password)
	 VALUES ($1, $2, $3, $4, $5)`,
		id, user.Name, user.Role, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserPostgres) DeleteUser(id string) error {
	_, err := u.db.Exec(`DELETE FROM users WHERE id = $1`, id)
	return err
}

func (u *UserPostgres) FindUserByID(id string) (models.User, error) {
	var user models.User

	err := u.db.QueryRow(`SELECT id, name, role, email FROM users WHERE id = $1`, id).
		Scan(&user.Id, &user.Name, &user.Role, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, nil
		}
		return models.User{}, err
	}

	return user, nil
}

func (u *UserPostgres) FindUserByEmail(email string) (models.User, error) {
	var user models.User

	err := u.db.QueryRow(`SELECT id, name, role, email, password FROM users WHERE email = $1`, email).
		Scan(&user.Id, &user.Name, &user.Role, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, nil
		}
		return models.User{}, err
	}

	return user, nil
}

// Need to implement
func (u *UserPostgres) UpdateUser(user models.UpdateUserDTO) error {
	query := `
		UPDATE users
		SET 
			name = $1, 
			role = $2, 
			email = $3
		WHERE id = $4
	`

	_, err := u.db.Exec(query, user.Name, user.Role, user.Email, user.Id)
	if err != nil {
		return err
	}
	return nil
}
