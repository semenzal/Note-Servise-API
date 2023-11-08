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

	type Note struct {
		id         int64
		title      string
		text       string
		author     string
		created_at time.Time
		updated_at sql.NullTime
	}
	Notes := []*desc.Note{}
	Notes = append(Notes, Notes...)

	for row.Next() {
		var id int64
		var title, text, author string
		var createdAt time.Time
		var updatedAt sql.NullTime
		err = row.Scan(&id, &title, &text, &author, &createdAt, &updatedAt)
		if err != nil {
			return nil, err
		}
	}

	return &desc.GetListResponse{
		Notes: []*desc.Note{},
	}, nil
}
