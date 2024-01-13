package app

import (
	"context"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/semenzal/note-service-api/internal/app/api/note_v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	grpcValidator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	desc "github.com/semenzal/note-service-api/pkg/note_v1"
	"google.golang.org/grpc"
)

// App ...
type App struct {
	noteImpl 			*note_v1.Note
	serviceProvider 	*serviceProvider

	pathConfig string

	grpcServer 	*grpc.Server
	mux 		*runtime.ServeMux
}

// NewApp ...
func NewApp(ctx context.Context, pathConfig string) (*App, error) {
	a := &App{
		pathConfig: pathConfig,
	}
	err := a.initDeps(ctx)

	return a, err
}

// Run ...
func (a *App) Run() error {
	defer func ()  {
		a.serviceProvider.db.Close()
	}()

	wg := &sync.WaitGroup{}
	wg.Add(2)

	err := a.runGRPC(wg)
	if err != nil {
		return err
	}

	err = a.runPublicHTTP(wg)
	if err != nil {
		return err
	}

	wg.Wait()
	return nil
}


func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error {
		a.initSreviceProvider,
		a.initServer,
		a.initGRPCServer,
		a.initPulicHTTPHandlers,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initSreviceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider(a.pathConfig)
	return nil
}

func (a *App) initServer (ctx context.Context) error {
	a.noteImpl = note_v1.NewNote(a.serviceProvider.GetNoteService(ctx))

	return nil
}

func (a *App) initGRPCServer(_ context.Context) error {
	a.grpcServer = grpc.NewServer(
		grpc.UnaryInterceptor(grpcValidator.UnaryServerInterceptor()),
	)

	desc.RegisterNoteServiceServer(a.grpcServer, a.noteImpl)

	return nil
}

func(a *App) initPulicHTTPHandlers(ctx context.Context) error {
	a.mux = runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := desc.RegisterNoteServiceHandlerFromEndpoint(ctx, a.mux, a.serviceProvider.GetConfig().GRPC.GetAddress(), opts)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) runGRPC(wg *sync.WaitGroup) error {
	list, err := net.Listen("tcp", a.serviceProvider.GetConfig().GRPC.GetAddress())
	if err != nil {
		return err
	}

	go func ()  {
		defer wg.Done()

		if err = a.grpcServer.Serve(list); err != nil {
			log.Fatal("failed to process gRPC server: %s", err.Error())
		}
	}()

	log.Printf("Run gRPC server on %s host\n", a.serviceProvider.GetConfig().GRPC.GetAddress())
	return nil
}

func (a *App) runPublicHTTP(wg *sync.WaitGroup) error {
	go func() {
		defer wg.Done()

		if err := http.ListenAndServe(a.serviceProvider.GetConfig().HTTP.GetAddress(), a.mux); err != nil {
			log.Fatalf("failed to process muxer: %s", err.Error())
		}
	}()

	log.Printf("Run public http handler on %s host\n", a.serviceProvider.GetConfig().HTTP.GetAddress())
	return nil
}