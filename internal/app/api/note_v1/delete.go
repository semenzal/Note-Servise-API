package note_v1

import (
	"context"
	"fmt"

	desc "github.com/semenzal/note-service-api/pkg/note_v1"
)

func (n *Note) DeleteNote(ctx context.Context, req *desc.DeleteRequest) (*desc.DeleteResponse, error) {
	fmt.Println("DeleteNote")
	fmt.Println("delete:", req.GetDelete())

	return &desc.DeleteResponse{
		Id: 0,
	}, nil
}
