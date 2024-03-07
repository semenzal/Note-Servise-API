package note_v1

import (
	"github.com/semenzal/note-service-api/internal/service/note"
	desc "github.com/semenzal/note-service-api/pkg/note_v1"
)

type Note struct {
	desc.UnimplementedNoteServiceServer

	noteService *note.Service
}

func NewNote(noteService *note.Service) *Note {
	return &Note{
		noteService: noteService,
	}
}

func newMockNoteV1(n Note) *Note {
	return &Note{
		desc.UnimplementedNoteServiceServer{},
		n.noteService,
	}
}
