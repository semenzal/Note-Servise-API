package note

import (
	"context"

	desc "github.com/semenzal/note-service-api/pkg/note_v1"
)

func (s *Service) Delete(ctx context.Context, req *desc.DeleteRequest) error {
	err := s.noteRepository.Delete(ctx, req)
	if err != nil {
		return err
	}

	return nil
}