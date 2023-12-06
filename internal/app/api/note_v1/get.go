package note_v1

import (
	"context"
	
	_ "github.com/jackc/pgx/stdlib"

	desc "github.com/semenzal/note-service-api/pkg/note_v1"
)

func (n *Note) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	res, err := n.noteService.Get(ctx, req)
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

	

	return &desc.GetResponse{
		Note: &desc.Note{
			Id:        id,
			Title:     title,
			Text:      text,
			Author:    author,
			CreatedAt: timestamppb.New(createdAt),
			UpdatedAt: updatedAtProto,
		},
	}, nil*/
}
