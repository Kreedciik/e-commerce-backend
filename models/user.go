package models

type User struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CreateUserDTO struct {
	Name     string `json:"name"`
	Role     string `json:"role"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserDTO struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Role  string `json:"role"`
	Email string `json:"email"`
}
