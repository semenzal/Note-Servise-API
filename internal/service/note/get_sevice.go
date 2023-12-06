package note

import (
	"context"
	desc "github.com/semenzal/note-service-api/pkg/note_v1"
)


func (s *Service) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	note, err := s.noteRepository.Get(ctx, req)
	if err != nil {
		return nil, err
	}

	return note, nil
}