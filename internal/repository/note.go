package repository

import (
	"context"
	"time"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/semenzal/note-service-api/internal/repository/table"
	desc "github.com/semenzal/note-service-api/pkg/note_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
	"github.com/golang/protobuf/ptypes/empty"

	sq "github.com/Masterminds/squirrel"
)

type NoteRepository interface {
	Create(ctx context.Context, req *desc.CreateRequest) (int64, error)
	Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error)
	GetList(ctx context.Context, req *empty.Empty) (*desc.GetListResponse, error)
	Update(ctx context.Context, req *desc.UpdateRequest) (*empty.Empty, error)
	Delete(ctx context.Context, req *desc.DeleteRequest) (*empty.Empty, error)
}

type repository struct {
	db *sqlx.DB
}

func NewNoteRepository(db *sqlx.DB) NoteRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(ctx context.Context, req *desc.CreateRequest) (int64, error){
	builder := sq.Insert(table.Note).
		PlaceholderFormat(sq.Dollar).
		Columns("title, text, author").
		Values(req.GetTitle(), req.GetText(), req.GetAuthor()).
		Suffix("returning id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	row, err := r.db.QueryContext(ctx, query, args...)
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

func (r *repository) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	builder := sq.Select("id", "title", "text", "author", "created_at", "updated_at").
		From(table.Note).
		Where(sq.Eq{"id": req.Id}).
		PlaceholderFormat(sq.Dollar).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	row := r.db.QueryRowContext(ctx, query, args...)

	var id int64
	var title, text, author string
	var createdAt time.Time
	var updatedAt sql.NullTime
	err = row.Scan(&id, &title, &text, &author, &createdAt, &updatedAt)
	if err != nil {
		return nil, err
	}

	var updatedAtProto *timestamppb.Timestamp
	if updatedAt.Valid {
		updatedAtProto = timestamppb.New(updatedAt.Time)
	}

	return &desc.GetResponse{
		Note: &desc.Note{
			Id:        id,
			Title:     title,
			Text:      text,
			Author:    author,
			CreatedAt: timestamppb.New(createdAt),
			UpdatedAt: updatedAtProto,
		},
	}, nil
}

func (r *repository) GetList(ctx context.Context, req *empty.Empty) (*desc.GetListResponse, error){
	builder := sq.Select("id", "title", "text", "author", "created_at", "updated_at").
		From(table.Note).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	row, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	notes := []*desc.Note{}

	for row.Next() {
		var id int64
		var title, text, author string
		var createdAt time.Time
		var updatedAt sql.NullTime
		err = row.Scan(&id, &title, &text, &author, &createdAt, &updatedAt)
		if err != nil {
			return nil, err
		}

		var updatedAtProto *timestamppb.Timestamp
		if updatedAt.Valid {
			updatedAtProto = timestamppb.New(updatedAt.Time)
		}

		notes = append(notes, &desc.Note{
			Id:        id,
			Title:     title,
			Text:      text,
			Author:    author,
			CreatedAt: timestamppb.New(createdAt),
			UpdatedAt: updatedAtProto,
		})
	}

	return &desc.GetListResponse{
		Notes: notes,
	}, nil
}

func (r *repository) Update(ctx context.Context, req *desc.UpdateRequest) (*empty.Empty, error) {
	builder := sq.Update(table.Note).
		Set("title", req.Title).
		Set("text", req.Text).
		Set("author", req.Author).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": req.Id})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	_, err = r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (r *repository) Delete(ctx context.Context, req *desc.DeleteRequest) (*empty.Empty, error) {
	builder := sq.Delete(table.Note).
		Where(sq.Eq{"id": req.Id}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	_, err = r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}