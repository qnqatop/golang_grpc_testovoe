package main

import (
	"fmt"
	"golang_grpc_testovoe/lib/directory"
	"golang_grpc_testovoe/server/config"
	"golang_grpc_testovoe/server/controller"
	"golang_grpc_testovoe/server/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	config.ApplicationConfig = config.NewConfig()

	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	directory.RegisterDirectoryInfoServiceServer(s, controller.NewDirectoryInfoController(service.NewDirectoryIndoService()))
	fmt.Println("Server started")

	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
