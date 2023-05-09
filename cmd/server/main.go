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
	userUsecase := usecase.NewUserUsecase(userRepository)
	userHandler := handler.NewUserHandler(userUsecase)

	server := server.NewServer(userHandler)
	server.Logger.Fatal(server.Start(":8080"))
}
