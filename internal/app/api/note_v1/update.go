package note_v1

import (
	"context"
	"fmt"

	desc "github.com/semenzal/note-service-api/pkg/note_v1"
)

func (n *Note) UpdateNote(ctx context.Context, req *desc.UpdateRequest) (*desc.UpdateResponse, error) {
	fmt.Println("UpdateNote")
	fmt.Println("update:", req.GetUpdate())

	return &desc.UpdateResponse{
		Id: 77,
	}, nil
}
