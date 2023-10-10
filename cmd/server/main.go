package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	desc "github.com/semenzal/note-service-api/pkg/note_v1"
)

const grpcPort = 50051

type server struct {
	desc.UnimplementedNoteServiceServer
}

// Create ...
func (s *server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf(
		"Title: %v\nAuthor:%v\nText:%v", req.GetTitle(), req.GetAuthor(), req.GetText(),
	)

	return &desc.CreateResponse{
		Id: int64(12),
	}, nil
}

func (s *server) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	log.Println(
		"Nota:", req.GetNota(),
	)

	return &desc.GetResponse{
		Id: int64(3),
	}, nil
}

func (s *server) GetList(ctx context.Context, req *desc.GetListRequest) (*desc.GetListResponse, error) {
	log.Println(
		"Nota:", req.GetAllId(),
	)

	return &desc.GetListResponse{
		Id: int64(21),
	}, nil
}

func (s *server) Update(ctx context.Context, req *desc.UpdateRequest) (*desc.UpdateResponse, error) {
	log.Println(
		"Update:", req.GetUpdate(),
	)

	return &desc.UpdateResponse{
		Id: int64(77),
	}, nil
}

func (s *server) Delete(ctx context.Context, req *desc.DeleteRequest) (*desc.DeleteResponse, error) {
	log.Println(
		"Delete:", req.GetDelete(),
	)

	return &desc.DeleteResponse{
		Id: int64(0),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterNoteServiceServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
