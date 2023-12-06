package note_v1

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	_ "github.com/jackc/pgx/stdlib"

	desc "github.com/semenzal/note-service-api/pkg/note_v1"
)

func (n *Note) GetList(ctx context.Context, req *empty.Empty) (*desc.GetListResponse, error) {
	res, err := n.noteService.GetList(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
	
	/*dbDsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, dbUser, dbPassword, dbName, sslMode,
	)

	db, err := sqlx.Open("pgx", dbDsn)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	

	return &desc.GetListResponse{
		Notes: notes,
	}, nil*/
}
