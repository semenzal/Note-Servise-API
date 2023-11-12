package note_v1

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/golang/protobuf/ptypes/empty"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"google.golang.org/protobuf/types/known/timestamppb"

	desc "github.com/semenzal/note-service-api/pkg/note_v1"
)

func (n *Note) GetList(ctx context.Context, req *empty.Empty) (*desc.GetListResponse, error) {

	dbDsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, dbUser, dbPassword, dbName, sslMode,
	)

	db, err := sqlx.Open("pgx", dbDsn)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	builder := sq.Select("id", "title", "text", "author", "created_at", "updated_at").
		From(noteTable).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	row, err := db.QueryContext(ctx, query, args...)
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
