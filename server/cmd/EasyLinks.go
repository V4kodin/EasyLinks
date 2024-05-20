package main

import (
	"EasyLinks/server"
	"EasyLinks/server/pkg/handler"
	"EasyLinks/server/pkg/service"
	"EasyLinks/server/storage"
	"log"
)

func main() {
	db, err := storage.InitDB()
	if err != nil {
		log.Fatalf("Error occured while connecting to database: %s", err.Error())
	}

	Service := service.NewService(db)
	Handlers := handler.NewHandler(Service)
	s := new(server.Server)

	port := server.Port
	if err := s.Start(port, Handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occured while running http server: %s", err.Error())
	} else {
		println("Server is running on port " + port)
	}
}
