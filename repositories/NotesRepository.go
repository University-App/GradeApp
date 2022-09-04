package repositories

import (
	"fmt"
	"github.com/paulmarie/univesity/grade_app/entities"
	"gorm.io/gorm"
)

type NoteRepository struct {
	DB *gorm.DB
}

func NewNoteRepository(db *gorm.DB) NoteRepository {
	return NoteRepository{db}
}

func (noteRepository NoteRepository) FindAllNotes() []entities.Note {
	var notes []entities.Note

	if result := noteRepository.DB.Find(&notes); result.Error != nil {
		fmt.Println(result.Error)
	}
	return notes
}

func (noteRepository NoteRepository) AddNote(note *entities.Note) entities.Note {

	if result := noteRepository.DB.Create(&note); result.Error != nil {
		fmt.Println(result.Error)
	}
	return *note
}
