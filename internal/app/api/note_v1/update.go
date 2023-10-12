package note_v1

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"

	desc "github.com/semenzal/note-service-api/pkg/note_v1"
)

func (n *Note) UpdateNote(ctx context.Context, req *desc.UpdateRequest) (*empty.Empty, error) {
	fmt.Println("UpdateNote")
	fmt.Println("update:", req.GetTitle(), req.GetText(), req.GetAuthor())

	return &empty.Empty{}, nil
}
