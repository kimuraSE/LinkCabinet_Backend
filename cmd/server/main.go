package main

import (
	"LinkCabinet_Backend/internal/api/handler"
	"LinkCabinet_Backend/internal/api/repository"
	"LinkCabinet_Backend/internal/api/usecase"
	"LinkCabinet_Backend/internal/api/validator"
	"LinkCabinet_Backend/internal/db"
	"LinkCabinet_Backend/pkg/server"
)

func main() {
	db := db.NewDB()

	userRepository := repository.NewUserRepository(db)
	linkRepository := repository.NewLinksRepository(db)

	userValidator := validator.NewUserValidator()
	linkValidator := validator.NewLinksValidator()

	userUsecase := usecase.NewUserUsecase(userRepository,userValidator)
	linkUsecase := usecase.NewLinksUsecase(linkRepository,linkValidator)

	userHandler := handler.NewUserHandler(userUsecase)
	linkHandler := handler.NewLinksHandler(linkUsecase)

	server := server.NewServer(userHandler,linkHandler)
	server.Logger.Fatal(server.Start(":8080"))
}
