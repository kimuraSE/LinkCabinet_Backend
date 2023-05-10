package main

import (
	"LinkCabinet_Backend/internal/api/handler"
	"LinkCabinet_Backend/internal/api/repository"
	"LinkCabinet_Backend/internal/api/usecase"
	"LinkCabinet_Backend/internal/db"
	"LinkCabinet_Backend/pkg/server"
)

func main() {
	db := db.NewDB()

	userRepository := repository.NewUserRepository(db)
	linkRepository := repository.NewLinksRepository(db)

	userUsecase := usecase.NewUserUsecase(userRepository)
	linkUsecase := usecase.NewLinksUsecase(linkRepository)

	userHandler := handler.NewUserHandler(userUsecase)
	linkHandler := handler.NewLinksHandler(linkUsecase)

	server := server.NewServer(userHandler,linkHandler)
	server.Logger.Fatal(server.Start(":8080"))
}
