package note_v1

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"

	_ "github.com/jackc/pgx/stdlib"

	desc "github.com/semenzal/note-service-api/pkg/note_v1"
)

func (n *Note) Update(ctx context.Context, req *desc.UpdateRequest) (*empty.Empty, error) {
	err := n.noteService.Update(ctx, req)
	if err != nil {
		return nil, err
	}

	return nil, err
}
