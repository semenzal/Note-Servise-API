package note_v1

import (
	"context"
	"fmt"

	desc "github.com/semenzal/note-service-api/pkg/note_v1"
)

func (n *Note) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	fmt.Println("GetNote")
	fmt.Println("nota", req.GetNota())

	return &desc.GetResponse{
		Id: 3,
	}, nil
}
