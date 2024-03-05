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

func TestGetList(t *testing.T) {
	var (
		ctx         = context.Background()
		mockCtrl    = gomock.NewController(t)
		id          = gofakeit.Int64()
		title       = gofakeit.BeerName()
		text        = gofakeit.BeerStyle()
		author      = gofakeit.Name()
		email       = gofakeit.Email()
		createdAt   = gofakeit.Date()
		updatedAt   = gofakeit.Date()
		limit       = gofakeit.Int64()
		offset      = gofakeit.Int64()
		repoErrText = gofakeit.Phrase()

		req = &desc.GetListRequest{
			Filter: &desc.Filter{
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
				Limit: &wrapperspb.Int64Value{
					Value: limit,
				},
				Offset: &wrapperspb.Int64Value{
					Value: offset,
				},
			},
		}

		repoReq = []*model.Note{{
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
		},
		}

		repoErr = errors.New(repoErrText)
	)

	noteMock := noteMocks.NewMockNoteRepository(mockCtrl)
	gomock.InOrder(
		noteMock.EXPECT().GetList(ctx, req).Return(repoReq, nil),
		noteMock.EXPECT().GetList(ctx, req).Return(nil, repoErr),
	)

	api := newMockNoteV1(Note{
		noteService: note.NewMockNoteService(noteMock),
	})

	t.Run("getList success test", func(t *testing.T) {
		res, err := api.GetList(ctx, req)
		require.Nil(t, err)
		require.Equal(t, repoReq, res)
	})

	t.Run("note repo getList err", func(t *testing.T) {
		_, err := api.GetList(ctx, req)
		require.NotNil(t, err)
		require.Equal(t, repoErrText, err.Error())
	})
}
