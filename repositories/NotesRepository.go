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

func (noteRepository NoteRepository) FindAllNotesOfStudent(studentID uint) []entities.Note {
	var notes []entities.Note

	if result := noteRepository.DB.Where("student_id = ?", studentID).Find(&notes); result.Error != nil {
		fmt.Println(result.Error)
	}

	return notes
}

func (noteRepository NoteRepository) FindAllStudents() []entities.Student {
	var students []entities.Student

	if result := noteRepository.DB.Find(&students); result.Error != nil {
		fmt.Println(result.Error)
	}
	return students
}

func (noteRepository NoteRepository) FindAllCourses() []entities.Course {
	var courses []entities.Course

	if result := noteRepository.DB.Find(&courses); result.Error != nil {
		fmt.Println(result.Error)
	}
	return courses
}

func (noteRepository NoteRepository) FindAllUnites() []entities.Unite {
	var unites []entities.Unite

	if result := noteRepository.DB.Find(&unites); result.Error != nil {
		fmt.Println(result.Error)
	}
	return unites
}

func (noteRepository NoteRepository) CreateStudentGlobalAverage(studentGlobalAverage entities.StudentGlobalAverage) {

	if result := noteRepository.DB.Create(&studentGlobalAverage); result.Error != nil {
		fmt.Println(result.Error)
	}
}

func (noteRepository NoteRepository) CreateStudentUniteAverage(studentUniteAverage entities.StudentUniteAverage) {

	if result := noteRepository.DB.Create(&studentUniteAverage); result.Error != nil {
		fmt.Println(result.Error)
	}
}

func (noteRepository NoteRepository) CreateStudentCourseAverage(studentCourseAverage entities.StudentCourseAverage) {

	if result := noteRepository.DB.Create(&studentCourseAverage); result.Error != nil {
		fmt.Println(result.Error)
	}
}

func (noteRepository NoteRepository) FindAllStudentGlobalAverage() []entities.StudentGlobalAverage {
	var studentGlobalAverages []entities.StudentGlobalAverage

	if result := noteRepository.DB.Find(&studentGlobalAverages); result.Error != nil {
		fmt.Println(result.Error)
	}
	return studentGlobalAverages
}

func (noteRepository NoteRepository) FindAllStudentUnitesAverage() []entities.StudentUniteAverage {
	var studentUniteAverages []entities.StudentUniteAverage

	if result := noteRepository.DB.Find(&studentUniteAverages); result.Error != nil {
		fmt.Println(result.Error)
	}
	return studentUniteAverages
}

func (noteRepository NoteRepository) FindAllStudentCoursesAverage() []entities.StudentCourseAverage {
	var studentCourseAverages []entities.StudentCourseAverage

	if result := noteRepository.DB.Find(&studentCourseAverages); result.Error != nil {
		fmt.Println(result.Error)
	}
	return studentCourseAverages
}
