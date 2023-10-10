package note_v1

import (
	"context"
	"fmt"

	desc "github.com/semenzal/note-service-api/pkg/note_v1"
)

func (n *Note) GetListNote(ctx context.Context, req *desc.GetListRequest) (*desc.GetListResponse, error) {
	fmt.Println("GetListNote")
	fmt.Println("all_id:", req.GetAllId())

	return &desc.GetListResponse{
		Id: 12_3,
	}, nil
}
