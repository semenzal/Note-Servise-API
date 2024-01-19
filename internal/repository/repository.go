package repository

import (
	"context"

	"github.com/semenzal/note-service-api/internal/model"
	desc "github.com/semenzal/note-service-api/pkg/note_v1"
)

type NoteRepository interface {
	Create(ctx context.Context, noteInfo *model.NoteInfo) (int64, error)
	Get(ctx context.Context, id int64) (*model.Note, error)
	GetList(ctx context.Context) ([]*model.Note, error)
	Update(ctx context.Context, req *desc.UpdateRequest) error
	Delete(ctx context.Context, req *desc.DeleteRequest) error
}