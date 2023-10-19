package note_v1

import (
	"context"
	"fmt"

	desc "github.com/semenzal/note-service-api/pkg/note_v1"
)

func (n *Note) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	fmt.Println("Get")
	fmt.Println("note", &desc.Note{})

	return &desc.GetResponse{
		Note: &desc.Note{},
	}, nil
}
