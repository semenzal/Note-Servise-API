package note_v1

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	desc "github.com/semenzal/note-service-api/pkg/note_v1"
)

func (n *Note) Delete(ctx context.Context, req *desc.DeleteRequest) (*empty.Empty, error) {

	dbDsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, dbUser, dbPassword, dbName, sslMode,
	)

	db, err := sqlx.Open("pgx", dbDsn)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	builder := sq.Delete(noteTable).
		Where(sq.Eq{"id": req.Id}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	_, err = db.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}
