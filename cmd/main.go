package main

import (
	"ecommerce"
	"ecommerce/handler"
	"ecommerce/repository"
	"ecommerce/service"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "root"
	dbName   = "ecommerce"
	sslMode  = "disable"
)

func main() {

	db, err := repository.NewPostgres(repository.Config{
		Host:     host,
		Dbname:   dbName,
		Password: password,
		Port:     port,
		User:     user,
		SSLMode:  sslMode,
	})

	if err != nil {
		panic(err)
	}
	repository := repository.NewRepository(db)
	services := service.NewService(repository)

	handler := handler.NewHandler(services)
	server := new(ecommerce.Server)

	if err := server.Run("8000", handler.InitRoutes()); err != nil {
		panic(err)
	}

}
