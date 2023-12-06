package note_v1

import (
	"context"
	_ "github.com/jackc/pgx/stdlib"
	desc "github.com/semenzal/note-service-api/pkg/note_v1"
)

func (n *Note) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	res, err := n.noteService.Create(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
