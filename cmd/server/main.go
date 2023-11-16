package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/semenzal/note-service-api/internal/app/api/note_v1"
	desc "github.com/semenzal/note-service-api/pkg/note_v1"
)

const grpcPort = 50051

func main() {
	list, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err.Error())
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterNoteServiceServer(s, note_v1.NewNote())

	fmt.Println("Server is running on port:", grpcPort)

	if err = s.Serve(list); err != nil {
		log.Fatalf("failed to serve: %s", err.Error())
	}
}
