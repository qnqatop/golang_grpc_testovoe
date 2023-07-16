package main

import (
	"fmt"
	"golang_grpc_testovoe/backend/controller"
	"golang_grpc_testovoe/backend/service"
	"golang_grpc_testovoe/lib/directory"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {

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
