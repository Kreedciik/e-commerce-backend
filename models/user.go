package models

type User struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	CreatedAt string `json:"-"`
	UpdatedAt string `json:"-"`
}

type CreateUserDTO struct {
	Name     string `json:"name" binding:"required"`
	Role     string `json:"role"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateUserDTO struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Role  string `json:"role"`
	Email string `json:"email"`
}
