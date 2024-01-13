package note

import (
	"context"

	"github.com/semenzal/note-service-api/internal/model"
)


func (s *Service) Get(ctx context.Context, id int64) (*model.Note, error) {
	res, err := s.noteRepository.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return res, nil
}