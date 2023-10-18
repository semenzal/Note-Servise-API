package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/golang/protobuf/ptypes/empty"

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
	log.Printf("Title: %v\nAuthor:%v\nText:%v", req.GetTitle(), req.GetAuthor(), req.GetText())

	return &desc.CreateResponse{
		Id: int64(12),
	}, nil
}

func (s *server) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	log.Println("Nota:", req.GetId())

	return &desc.GetResponse{
		Nota: req.GetId(),
	}, nil
}

func (s *server) GetList(ctx context.Context, req *empty.Empty) (*desc.GetListResponse, error) {
	log.Println("Nota:")

	return &desc.GetListResponse{
		Notes: []*desc.Note{},
	}, nil
}

func (s *server) Update(ctx context.Context, req *desc.UpdateRequest) (*empty.Empty, error) {
	log.Println(
		"Title:", req.GetTitle(),
		"Text:", req.GetText(),
		"Author", req.GetAuthor(),
	)

	return &empty.Empty{}, nil
}

func (s *server) Delete(ctx context.Context, req *desc.DeleteRequest) (*empty.Empty, error) {
	log.Println(
		"Delete:", req.GetId(),
	)

	return &empty.Empty{}, nil
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
