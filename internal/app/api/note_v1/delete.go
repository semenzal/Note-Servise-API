package note_v1

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	_ "github.com/jackc/pgx/stdlib"

	desc "github.com/semenzal/note-service-api/pkg/note_v1"
)

func (n *Note) Delete(ctx context.Context, req *desc.DeleteRequest) (*empty.Empty, error) {
	err := n.noteService.Delete(ctx, req)
	if err != nil {
		return nil, err
	}

	return nil, err
}
