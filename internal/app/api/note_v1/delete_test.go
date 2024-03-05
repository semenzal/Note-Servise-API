package note_v1

import (
	"context"
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	noteMocks "github.com/semenzal/note-service-api/internal/repository/mocks"
	"github.com/semenzal/note-service-api/internal/service/note"
	desc "github.com/semenzal/note-service-api/pkg/note_v1"
	"github.com/stretchr/testify/require"
)

func TestDelete(t *testing.T) {
	var (
		ctx      = context.Background()
		mockCtrl = gomock.NewController(t)

		id          = gofakeit.Int64()
		repoErrText = gofakeit.Phrase()

		req = &desc.DeleteRequest{
			Id: id,
		}

		repoErr = errors.New(repoErrText)
	)

	noteMock := noteMocks.NewMockNoteRepository(mockCtrl)
	gomock.InOrder(
		noteMock.EXPECT().Create(ctx, id).Return(id, nil),
		noteMock.EXPECT().Create(ctx, id).Return(int64(0), repoErr),
	)

	api := newMockNoteV1(Note{
		noteService: note.NewMockNoteService(noteMock),
	})

	t.Run("delete success case", func(t *testing.T) {
		res, err := api.Delete(ctx, req)
		require.Nil(t, err)
		require.Equal(t, req.GetId(), res)
	})

	t.Run("note repo delete err", func(t *testing.T) {
		_, err := api.Delete(ctx, req)
		require.NotNil(t, err)
		require.Equal(t, repoErrText, err.Error())
	})
}
