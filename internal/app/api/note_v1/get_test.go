package note_v1

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/semenzal/note-service-api/internal/model"
	noteMocks "github.com/semenzal/note-service-api/internal/repository/mocks"
	"github.com/semenzal/note-service-api/internal/service/note"
	desc "github.com/semenzal/note-service-api/pkg/note_v1"
	"github.com/stretchr/testify/require"
)

func TestGet(t *testing.T) {
	var (
		ctx      = context.Background()
		mockCtrl = gomock.NewController(t)

		id          = gofakeit.Int64()
		title       = gofakeit.BeerName()
		text        = gofakeit.BeerStyle()
		author      = gofakeit.Name()
		email       = gofakeit.Email()
		createdAt   = gofakeit.Date()
		updatedAt   = gofakeit.Date()
		repoErrText = gofakeit.Phrase()

		req = &desc.GetRequest{
			Id: id,
		}

		validRes = &model.Note{
			ID: id,
			Info: &model.NoteInfo{
				Title:  title,
				Text:   text,
				Author: author,
				Email:  email,
			},
			CreatedAt: createdAt,
			UpdatedAt: sql.NullTime{
				Time:  updatedAt,
				Valid: true,
			},
		}

		repoErr = errors.New(repoErrText)
	)

	noteMock := noteMocks.NewMockNoteRepository(mockCtrl)
	gomock.InOrder(
		noteMock.EXPECT().Get(ctx, id).Return(id, nil),
		noteMock.EXPECT().Get(ctx, id).Return(nil, repoErr),
	)

	api := newMockNoteV1(Note{
		noteService: note.NewMockNoteService(noteMock),
	})

	t.Run("get success test", func(t *testing.T) {
		res, err := api.Get(ctx, req)
		require.Nil(t, err)
		require.Equal(t, validRes, res)
	})

	t.Run("note repo get err", func(t *testing.T) {
		_, err := api.Get(ctx, req)
		require.NotNil(t, err)
		require.Equal(t, repoErrText, err.Error())
	})
}
