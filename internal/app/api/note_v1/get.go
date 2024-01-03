package note_v1

import (
	"context"
	
	_ "github.com/jackc/pgx/stdlib"

	desc "github.com/semenzal/note-service-api/pkg/note_v1"
)

func (n *Note) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	res, err := n.noteService.Get(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
