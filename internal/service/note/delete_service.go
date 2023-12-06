package note

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	desc "github.com/semenzal/note-service-api/pkg/note_v1"
)

func (s *Service) Delete(ctx context.Context, req *desc.DeleteRequest) (*empty.Empty, error) {
	_, err := s.noteRepository.Delete(ctx, req)
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}