package note

import (
	"context"

	"github.com/semenzal/note-service-api/internal/model"
)

func (s *Service) Create(ctx context.Context, noteInfo *model.NoteInfo) (int64, error) {
	id, err := s.noteRepository.Create(ctx, noteInfo)
	if err != nil {
		return 0, err
	}

	return id, nil
}