package converter

import (
	"github.com/semenzal/note-service-api/internal/model"
	desc "github.com/semenzal/note-service-api/pkg/note_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToNoteInfo(noteInfo *desc.NoteInfo) *model.NoteInfo {
	return &model.NoteInfo{
		Title:  noteInfo.GetTitle(),
		Text:   noteInfo.GetText(),
		Author: noteInfo.GetAuthor(),
	}
}

func ToDescNoteInfo(noteInfo *model.NoteInfo) *desc.NoteInfo {
	return &desc.NoteInfo{
		Title:  noteInfo.Title,
		Text:   noteInfo.Text,
		Author: noteInfo.Author,
		Email: 	noteInfo.Email,
	}
}

func ToDescNote(note *model.Note) *desc.Note {
	var updatedAt *timestamppb.Timestamp
	if note.UpdatedAt.Valid {
		updatedAt = timestamppb.New(note.UpdatedAt.Time)
	}

	return &desc.Note{
		Id:        	note.ID,
		Info:      	ToDescNoteInfo(note.Info),
		CreatedAt: 	timestamppb.New(note.CreatedAt),
		UpdatedAt: 	updatedAt,
		Email: 		note.Info.Email,
	}
}

func ToDescNotes(notes []*model.Note) []*desc.Note {

	res := make([]*desc.Note, 0, len(notes))

	for _, elem := range notes {
		res = append(res, ToDescNote(elem))
	}

	return res
}
