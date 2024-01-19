package repository

import (
	"context"
	"time"

	"github.com/semenzal/note-service-api/internal/model"
	"github.com/semenzal/note-service-api/internal/pkg/db"
	"github.com/semenzal/note-service-api/internal/repository/table"
	desc "github.com/semenzal/note-service-api/pkg/note_v1"

	sq "github.com/Masterminds/squirrel"
)

type NoteRepository interface {
	Create(ctx context.Context, noteInfo *model.NoteInfo) (int64, error)
	Get(ctx context.Context, id int64) (*model.Note, error)
	GetList(ctx context.Context) ([]*model.Note, error)
	Update(ctx context.Context, req *desc.UpdateRequest) error
	Delete(ctx context.Context, req *desc.DeleteRequest) error
}

type repository struct {
	client db.Client
}

func NewNoteRepository(client db.Client) NoteRepository {
	return &repository{
		client: client,
	}
}

func (r *repository) Create(ctx context.Context, noteInfo *model.NoteInfo) (int64, error) {
	builder := sq.Insert(table.Note).
		PlaceholderFormat(sq.Dollar).
		Columns("title, text, author").
		Values(noteInfo.Title, noteInfo.Text, noteInfo.Author).
		Suffix("returning id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "Create",
		QueryRaw: query,
	}

	row, err := r.client.DB().QueryContext(ctx, q, args...)
	if err != nil {
		return 0, err
	}
	defer row.Close()

	row.Next()
	var id int64
	err = row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repository) Get(ctx context.Context, id int64) (*model.Note, error) {
	builder := sq.Select("id", "title", "text", "author", "created_at", "updated_at", "email").
		From(table.Note).
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "Get",
		QueryRaw: query,
	}

	note := new(model.Note)
	err = r.client.DB().GetContext(ctx, note, q, args...)
	if err != nil {
		return nil, err
	}

	return note, nil
}

func (r *repository) GetList(ctx context.Context) ([]*model.Note, error) {
	builder := sq.Select("id", "title", "text", "author", "created_at", "updated_at", "email").
		From(table.Note).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "GetList",
		QueryRaw: query,
	}

	var notes []*model.Note
	err = r.client.DB().SelectContext(ctx, notes, q, args...)
	if err != nil {
		return nil, err
	}

	return notes, nil
}

func (r *repository) Update(ctx context.Context, req *desc.UpdateRequest) error {
	builder := sq.Update(table.Note).
		PlaceholderFormat(sq.Dollar).
		Set("update_at", time.Now()).
		Where(sq.Eq{"id": req.GetId()})

	if req.GetNote().GetTitle() != nil {
		builder.Set("title", req.GetNote().GetTitle())
	}

	if req.GetNote().GetText() != nil {
		builder.Set("text", req.GetNote().GetText())
	}

	if req.GetNote().GetAuthor() != nil {
		builder.Set("author", req.GetNote().GetAuthor())
	}

	if req.GetNote().GetEmail() != nil {
		builder.Set("email", req.GetNote().GetEmail())
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "Update",
		QueryRaw: query,
	}

	_, err = r.client.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) Delete(ctx context.Context, req *desc.DeleteRequest) error {
	builder := sq.Delete(table.Note).
		Where(sq.Eq{"id": req.Id}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "Delete",
		QueryRaw: query,
	}
	_, err = r.client.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}
