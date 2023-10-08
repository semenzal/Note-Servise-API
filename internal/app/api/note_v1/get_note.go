package note_v1

import (
	"context"
	"fmt"

	desc "github.com/semenzal/note-service-api/pkg/note_v1"
)

func (n *Note) GetNote(ctx context.Context, req *desc.GetNoteRequest) (*desc.GetNoteResponse, error) {
	fmt.Println("GetNote")
	fmt.Println("nota", req.GetNota())

	return &desc.GetNoteResponse{
		Id: 3,
	}, nil
}
