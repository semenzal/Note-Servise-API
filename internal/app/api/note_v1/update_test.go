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
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestUpdate(t *testing.T) {
	var (
		ctx         = context.Background()
		mockCtrl    = gomock.NewController(t)
		id          = gofakeit.Int64()
		title       = gofakeit.BeerName()
		text        = gofakeit.BeerStyle()
		author      = gofakeit.Name()
		email       = gofakeit.Email()
		repoErrText = gofakeit.Phrase()

		req = &desc.UpdateRequest{
			Id: id,
			Note: &desc.UpdateNoteInfo{
				Title: &wrapperspb.StringValue{
					Value: title,
				},
				Text: &wrapperspb.StringValue{
					Value: text,
				},
				Author: &wrapperspb.StringValue{
					Value: author,
				},
				Email: &wrapperspb.StringValue{
					Value: email,
				},
			},
		}

		repoReq = &model.UpdateNoteInfo{
			Title: sql.NullString{
				String: title,
			},
			Text: sql.NullString{
				String: text,
			},
			Author: sql.NullString{
				String: author,
			},
			Email: sql.NullString{
				String: email,
			},
		}

		repoErr = errors.New(repoErrText)
	)

	noteMock := noteMocks.NewMockNoteRepository(mockCtrl)
	gomock.InOrder(
		noteMock.EXPECT().Update(ctx, id, repoReq).Return(nil),
		noteMock.EXPECT().Update(ctx, id, repoReq).Return(repoErr),
	)

	api := newMockNoteV1(Note{
		noteService: note.NewMockNoteService(noteMock),
	})

	t.Run("update success test", func(t *testing.T) {
		res, err := api.Update(ctx, req)
		require.Nil(t, err)
		require.Equal(t, repoReq, res)
	})

	t.Run("note repo update err", func(t *testing.T) {
		_, err := api.Update(ctx, req)
		require.NotNil(t, err)
		require.Equal(t, repoErrText, err.Error())
	})
}
