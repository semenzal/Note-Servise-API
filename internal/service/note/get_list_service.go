package note

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	desc "github.com/semenzal/note-service-api/pkg/note_v1"
)

func (s *Service) GetList(ctx context.Context, req *empty.Empty) (*desc.GetListResponse, error) {
	notes, err := s.noteRepository.GetList(ctx, req)
	if err != nil {
		return nil, err
	}

	return notes, nil
}