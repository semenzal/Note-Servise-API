package note_v1

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/semenzal/note-service-api/internal/converter"
	desc "github.com/semenzal/note-service-api/pkg/note_v1"
)

func (n *Note) Update(ctx context.Context, req *desc.UpdateRequest) (*empty.Empty, error) {
	err := n.noteService.Update(ctx, req.GetId(), converter.ToUpdateInfo(req.GetNote()))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
