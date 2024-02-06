package note

import (
	"context"

	"github.com/semenzal/note-service-api/internal/model"
)

func (s *Service) Update(ctx context.Context, id int64, updateInfo *model.UpdateNoteInfo) error {
	err := s.noteRepository.Update(ctx, id, updateInfo)
	if err != nil {
		return err
	}

	return nil
}
