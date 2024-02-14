package converter

import (
	"database/sql"

	"github.com/semenzal/note-service-api/internal/model"
	desc "github.com/semenzal/note-service-api/pkg/note_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToNoteInfo(noteInfo *desc.NoteInfo) *model.NoteInfo {
	return &model.NoteInfo{
		Title:  noteInfo.GetTitle(),
		Text:   noteInfo.GetText(),
		Author: noteInfo.GetAuthor(),
		Email:  noteInfo.GetEmail(),
	}
}

func ToUpdateInfo(updateInfo *desc.UpdateNoteInfo) *model.UpdateNoteInfo {
	var title, text, author, email sql.NullString

	if updateInfo.Title != nil {
		title = sql.NullString{
			String: updateInfo.Title.Value,
			Valid:  true,
		}
	}

	if updateInfo.Text != nil {
		text = sql.NullString{
			String: updateInfo.Text.Value,
			Valid:  true,
		}
	}

	if updateInfo.Author != nil {
		author = sql.NullString{
			String: updateInfo.Author.Value,
			Valid:  true,
		}
	}

	if updateInfo.Email != nil {
		email = sql.NullString{
			String: updateInfo.Email.Value,
			Valid:  true,
		}
	}

	return &model.UpdateNoteInfo{
		Title:  title,
		Text:   text,
		Author: author,
		Email:  email,
	}
}

func ToDescNoteInfo(noteInfo *model.NoteInfo) *desc.NoteInfo {
	return &desc.NoteInfo{
		Title:  noteInfo.Title,
		Text:   noteInfo.Text,
		Author: noteInfo.Author,
		Email:  noteInfo.Email,
	}
}

func ToDescNote(note *model.Note) *desc.Note {
	var updatedAt *timestamppb.Timestamp
	if note.UpdatedAt.Valid {
		updatedAt = timestamppb.New(note.UpdatedAt.Time)
	}

	return &desc.Note{
		Id:        note.ID,
		Info:      ToDescNoteInfo(note.Info),
		CreatedAt: timestamppb.New(note.CreatedAt),
		UpdatedAt: updatedAt,
	}
}

func ToDescNotes(notes []*model.Note) []*desc.Note {

	res := make([]*desc.Note, 0, len(notes))

	for _, elem := range notes {
		res = append(res, ToDescNote(elem))
	}

	return res
}

func ToFilterInfo(filter *desc.Filter) *model.Filter {
	var title, text, author, email sql.NullString
	var limit, offset sql.NullInt64

	if filter.Title != nil {
		title = sql.NullString{
			String: filter.Title.Value,
			Valid:  true,
		}
	}

	if filter.Text != nil {
		text = sql.NullString{
			String: filter.Text.Value,
			Valid:  true,
		}
	}

	if filter.Author != nil {
		author = sql.NullString{
			String: filter.Author.Value,
			Valid:  true,
		}
	}

	if filter.Email != nil {
		email = sql.NullString{
			String: filter.Email.Value,
			Valid:  true,
		}
	}

	if filter.Limit != nil {
		limit = sql.NullInt64{
			Int64: filter.Limit.Value,
			Valid: true,
		}
	}

	if filter.Offset != nil {
		offset = sql.NullInt64{
			Int64: filter.Offset.Value,
			Valid: true,
		}
	}

	return &model.Filter{
		Title:  title,
		Text:   text,
		Author: author,
		Email:  email,
		Limit:  limit,
		Offset: offset,
	}
}
