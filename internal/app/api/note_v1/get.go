package note_v1

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"google.golang.org/protobuf/types/known/timestamppb"

	desc "github.com/semenzal/note-service-api/pkg/note_v1"
)

func (n *Note) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
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
		Where(sq.Eq{"id": req.Id}).
		PlaceholderFormat(sq.Dollar).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	row := db.QueryRowContext(ctx, query, args...)

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
