package note_v1

import (
	"context"
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

func TestCreate(t *testing.T) {
	var (
		ctx      = context.Background()
		mockCtrl = gomock.NewController(t)

		id          = gofakeit.Int64()
		title       = gofakeit.BeerName()
		text        = gofakeit.BeerStyle()
		author      = gofakeit.Name()
		repoErrText = gofakeit.Phrase()

		req = &desc.CreateRequest{
			Note: &desc.NoteInfo{
				Title:  title,
				Text:   text,
				Author: author,
			},
		}

		repoReq = &model.NoteInfo{
			Title:  title,
			Text:   text,
			Author: author,
		}

		validRes = &desc.CreateResponse{
			Id: id,
		}

		repoErr = errors.New(repoErrText)
	)

	noteMock := noteMocks.NewMockNoteRepository(mockCtrl)
	gomock.InOrder(
		noteMock.EXPECT().Create(ctx, repoReq).Return(id, nil),
		noteMock.EXPECT().Create(ctx, repoReq).Return(int64(0), repoErr),
	)

	api := newMockNoteV1(Note{
		noteService: note.NewMockNoteService(noteMock),
	})

	t.Run("create success case", func(t *testing.T) {
		res, err := api.Create(ctx, req)
		require.Nil(t, err)
		require.Equal(t, validRes, res)
	})

	t.Run("note repo create err", func(t *testing.T) {
		_, err := api.Create(ctx, req)
		require.NotNil(t, err)
		require.Equal(t, repoErrText, err.Error())
	})
}
