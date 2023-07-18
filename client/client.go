package main

import (
	"context"
	"fmt"
	"golang_grpc_testovoe/lib/directory"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := directory.NewDirectoryInfoServiceClient(conn)

	// Пример запроса к серверу
	req := &directory.DirectoryRequest{
		// указать путь до папки
		Path: "",
	}

	response, err := client.GetInformationForPath(context.Background(), req)
	if err != nil {
		log.Fatalf("Error calling GetDirectory: %v", err)
	}
	var a string
	for _, dir := range response.Directories {
		a += fmt.Sprintf("dirName - %s,dirSize - %s", dir.Name, dir.Size) + "\n"
	}
	for _, file := range response.Files {
		a += fmt.Sprintf("fileName - %s,fileSize - %s", file.Name, file.Size) + "\n"
	}
	fmt.Print(a)
	a = ""

	// Допустим, через некоторое время делается повторный запрос к тому же пути
	time.Sleep(5 * time.Second)

	response, err = client.GetInformationForPath(context.Background(), req)
	if err != nil {
		log.Fatalf("Error calling GetDirectory: %v", err)
	}
	for _, dir := range response.Directories {
		a += fmt.Sprintf("dirName - %s,dirSize - %s", dir.Name, dir.Size) + "\n"
	}
	for _, file := range response.Files {
		a += fmt.Sprintf("fileName - %s,fileSize - %s", file.Name, file.Size) + "\n"
	}
	fmt.Print(a)
}
