package note_v1

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"

	desc "github.com/semenzal/note-service-api/pkg/note_v1"
)

func (n *Note) GetList(ctx context.Context, req *empty.Empty) (*desc.GetListResponse, error) {
	fmt.Println("GetList")
	fmt.Println("all_id:", req.String())

	return &desc.GetListResponse{
		Notes: []*desc.Note{},
	}, nil
}
