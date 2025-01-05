package service

import (
	"ecommerce/config"
	"ecommerce/models"
	"ecommerce/repository"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repository repository.User
}

// All methods of user service should be implemented here
type User interface {
	CreateUser(user models.CreateUserDTO) error
	UpdateUser(user models.UpdateUserDTO) error
	DeleteUser(id string) error
	GetUserByID(id string) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
	HashPassword(password string) (string, error)
	ComparePasswords(hashedPassword string, password string) error
	GenerateAccessToken(user models.User, ttl time.Duration, secretKey string) (string, error)
	ParseToken(tokenString string) (*models.UserClaims, error)
}

func NewUserService(repository repository.User) *UserService {
	return &UserService{
		repository,
	}
}

func (u *UserService) CreateUser(user models.CreateUserDTO) error {
	hashedPassword, err := u.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	return u.repository.InsertUser(user)
}

func (u *UserService) UpdateUser(user models.UpdateUserDTO) error {
	return u.repository.UpdateUser(user)
}

func (u *UserService) DeleteUser(id string) error {
	return u.repository.DeleteUser(id)
}

func (u *UserService) GetUserByID(id string) (models.User, error) {
	return u.repository.FindUserByID(id)
}

func (u *UserService) GetUserByEmail(email string) (models.User, error) {
	return u.repository.FindUserByEmail(email)
}

func (u *UserService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (u *UserService) ComparePasswords(hashedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *UserService) GenerateAccessToken(user models.User, ttl time.Duration, secretKey string) (string, error) {

	claims := models.UserClaims{
		user.Role,
		user.Id,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttl)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secretKey))
	return tokenString, err
}

func (u *UserService) ParseToken(tokenString string) (*models.UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return models.UserClaims{}, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.SECRET_KEY), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(*models.UserClaims)

	if ok {
		if claims.ExpiresAt.Unix() < time.Now().Unix() {
			return nil, errors.New("token has expired")
		}

		if claims.UserId == "" {
			return nil, errors.New("user ID is missing in token")
		}
	}
	return claims, nil
}
