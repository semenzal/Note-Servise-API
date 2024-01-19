package note_v1

import (
	"context"

	"github.com/semenzal/note-service-api/internal/converter"
	desc "github.com/semenzal/note-service-api/pkg/note_v1"
)

func (n *Note) GetList(ctx context.Context) (*desc.GetListResponse, error) {
	notes, err := n.noteService.GetList(ctx)
	if err != nil {
		return nil, err
	}

	return &desc.GetListResponse{
		Notes: converter.ToDescNotes(notes),
	}, nil
}
