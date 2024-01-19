package note_v1

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/semenzal/note-service-api/pkg/note_v1"
)

func (n *Note) Delete(ctx context.Context, req *desc.DeleteRequest) (*empty.Empty, error) {
	err := n.noteService.Delete(ctx, req)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
