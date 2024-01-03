package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"

	grpcValidator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/semenzal/note-service-api/internal/app/api/note_v1"
	"github.com/semenzal/note-service-api/internal/repository"
	"github.com/semenzal/note-service-api/internal/service/note"
	desc "github.com/semenzal/note-service-api/pkg/note_v1"
)

const (
	hostGrpc = "localhost:50051"
	hostHttp = "localhost:8090"
)

const (
	host       = "localhost"
	port       = "54321"
	dbUser     = "note-service-user"
	dbPassword = "note-service-password"
	dbName     = "note-service"
	sslMode    = "disable"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	
	go func() {
		defer wg.Done()
		err := starGRPC()
		if err != nil {
			log.Fatal(err.Error())
		}
	}()

	go func() {
		defer wg.Done()
		err := startHttp()
		if err != nil {
			log.Fatal(err.Error())
		}
	}()

	wg.Wait()
}

func starGRPC() error {
	list, err := net.Listen("tcp", hostGrpc)
	if err != nil {
		return err
	}

	dbDsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, dbUser, dbPassword, dbName, sslMode,
	)

	db, err := sqlx.Open("pgx", dbDsn)
	if err != nil {
		return err
	}
	defer db.Close()

	noteRepository := repository.NewNoteRepository(db)
	noteService := note.NewService(noteRepository)

	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpcValidator.UnaryServerInterceptor()),
	)
	reflection.Register(s)
	desc.RegisterNoteServiceServer(s, note_v1.NewNote(noteService))

	fmt.Println("grpc server is running on port:", hostGrpc)

	if err = s.Serve(list); err != nil {
		return err
	}

	return nil
}

func startHttp() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := desc.RegisterNoteServiceHandlerFromEndpoint(ctx, mux, hostGrpc, opts)
	if err != nil {
		return err
	}

	fmt.Println("http server is running on port:", hostHttp)
	return http.ListenAndServe(hostHttp, mux)
}
