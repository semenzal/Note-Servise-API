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

func (s *server) GetNote(ctx context.Context, req *desc.GetNoteRequest) (*desc.GetNoteResponse, error) {
	log.Println(
		"Nota:", req.GetNota(),
	)

	return &desc.GetNoteResponse{
		Id: int64(3),
	}, nil
}

func (s *server) GetListNote(ctx context.Context, req *desc.GetListNoteRequest) (*desc.GetListNoteResponse, error) {
	log.Println(
		"Nota:", req.GetAllId(),
	)

	return &desc.GetListNoteResponse{
		Id: int64(21),
	}, nil
}

func (s *server) UpdateNote(ctx context.Context, req *desc.UpdateNoteRequest) (*desc.UpdateNoteResponse, error) {
	log.Println(
		"Update:", req.GetUpdate(),
	)

	return &desc.UpdateNoteResponse{
		Id: int64(77),
	}, nil
}

func (s *server) DeleteNote(ctx context.Context, req *desc.DeleteNoteRequest) (*desc.DeleteNoteResponse, error) {
	log.Println(
		"Delete:", req.GetDelete(),
	)

	return &desc.DeleteNoteResponse{
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
