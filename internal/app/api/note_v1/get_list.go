package note_v1

import (
	"context"
	"fmt"

	/*"go/build"*/

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

	builder := sq.Select("id").
		From(noteTable).
		Where(sq.Eq{"Notes": []*desc.Note{}}).
		PlaceholderFormat(sq.Dollar).
		Suffix("return all id")

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	row, err := db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	row.Next()
	var notes Note
	err = row.Scan(&notes)
	if err != nil {
		return nil, err
	}

	return &desc.GetListResponse{
		Notes: []*desc.Note{},
	}, nil
}
