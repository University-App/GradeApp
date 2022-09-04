package services

import (
	"github.com/paulmarie/univesity/grade_app/entities"
	"github.com/paulmarie/univesity/grade_app/repositories"
	"gorm.io/gorm"
)

type NoteService struct {
	reposiroty repositories.NoteRepository
}

func NewNoteService(db *gorm.DB) NoteService {
	return NoteService{repositories.NewNoteRepository(db)}
}

func (noteService NoteService) GetAllNotes() []entities.Note {
	return noteService.reposiroty.FindAllNotes()
}

func (noteService NoteService) AddNote(note *entities.Note) entities.Note {

	return noteService.reposiroty.AddNote(note)
}

func (noteService NoteService) ComputeAverage() int {

	average := 0

	notes := noteService.GetAllNotes()

	for _, note := range notes {
		average += note.Nombre
	}
	return average / len(notes)
}
