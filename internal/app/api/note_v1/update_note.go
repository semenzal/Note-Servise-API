package note_v1

import (
	"context"
	"fmt"

	desc "github.com/semenzal/note-service-api/pkg/note_v1"
)

func (n *Note) UpdateNote(ctx context.Context, req *desc.UpdateNoteRequest) (*desc.UpdateNoteResponse, error) {
	fmt.Println("UpdateNote")
	fmt.Println("update:", req.GetUpdate())

	return &desc.UpdateNoteResponse{
		Id: 77,
	}, nil
}
