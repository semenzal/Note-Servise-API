package note_v1

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"

	desc "github.com/semenzal/note-service-api/pkg/note_v1"
)

func (n *Note) DeleteNote(ctx context.Context, req *desc.DeleteRequest) (*empty.Empty, error) {
	fmt.Println("DeleteNote")
	fmt.Println("delete:", req.GetId())

	return &empty.Empty{}, nil
}
