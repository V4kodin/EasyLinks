package main

import (
	"EasyLinks/server"
	"EasyLinks/server/pkg/handler"
	"EasyLinks/server/pkg/service"
	"EasyLinks/server/storage"
	"log"
)

const (
	port          = ":8080"
	serverAddress = "http://localhost"
)

func main() {
	db, err := storage.InitDB()
	if err != nil {

	}
	service := service.NewService(db)
	handlers := handler.NewHandler(service)
	s := new(server.Server)

	if err := s.Start(port, handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occured while running http server: %s", err.Error())
	} else {
		println("Server is running on port 8080")
	}
}
