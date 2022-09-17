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

func (noteService NoteService) AddNote(note *entities.Note) entities.Note {

	return noteService.reposiroty.AddNote(note)
}

func (noteService NoteService) ComputeGlobalAverageForEachStudent() []entities.StudentAverage {
	var studentsGlobalAverages []entities.StudentAverage

	students := noteService.reposiroty.FindAllStudents()

	for _, student := range students {
		gobalAverageResult := 0
		var notes []entities.Note
		var globalAverage entities.StudentAverage
		notes = noteService.reposiroty.FindAllNotesOfStudent(student.ID)
		if len(notes) > 0 {
			for _, note := range notes {
				gobalAverageResult += note.Nombre
			}

			globalAverage.Average = gobalAverageResult / len(notes)
		}
		globalAverage.StudentName = student.FirstName + "-" + student.LastName

		studentsGlobalAverages = append(studentsGlobalAverages, globalAverage)
	}

	return studentsGlobalAverages
}

func (noteService NoteService) ComputeUniteAverageForEachStudent() []entities.StudentUniteAverage {
	var studentUniteAverages []entities.StudentUniteAverage

	unites := noteService.reposiroty.FindAllUnites()
	courses := noteService.reposiroty.FindAllCourses()
	students := noteService.reposiroty.FindAllStudents()

	for _, student := range students {
		for _, unite := range unites {
			var notes []entities.Note
			var courseAverage entities.StudentUniteAverage
			notes = noteService.reposiroty.FindAllNotesOfStudent(student.ID)
			courseAverageResult := 0
			compteur := 0
			if len(notes) > 0 {
				for _, course := range courses {
					if course.UniteID == unite.ID {
						for _, note := range notes {
							if note.CourseName == course.Name {
								courseAverageResult += note.Nombre
								compteur++
							}
						}
					}
				}
				courseAverage.StudentAverage.Average = courseAverageResult / compteur
				courseAverage.UniteName = unite.Name
				courseAverage.StudentAverage.StudentName = student.FirstName + "-" + student.LastName
				studentUniteAverages = append(studentUniteAverages, courseAverage)
			}
		}
	}
	return studentUniteAverages
}

func (noteService NoteService) ComputeCourseAverageForEachStudent() []entities.StudentCourseAverage {
	var studentCourseAverages []entities.StudentCourseAverage

	students := noteService.reposiroty.FindAllStudents()
	courses := noteService.reposiroty.FindAllCourses()

	for _, student := range students {
		var notes []entities.Note
		var courseAverage entities.StudentCourseAverage
		notes = noteService.reposiroty.FindAllNotesOfStudent(student.ID)
		if len(notes) > 0 {
			for _, course := range courses {
				courseAverageResult := 0
				compteur := 0
				for _, note := range notes {
					if note.CourseName == course.Name {
						courseAverageResult += note.Nombre
						compteur++
					}
				}
				courseAverage.StudentAverage.Average = courseAverageResult / compteur
				courseAverage.StudentAverage.StudentName = student.FirstName + "-" + student.LastName
				courseAverage.CourseName = course.Name
				studentCourseAverages = append(studentCourseAverages, courseAverage)
			}
		}
	}
	return studentCourseAverages
}

func (noteService NoteService) ComputeGlobalAverage() entities.GlobalAverage {

	average := 0
	notes := noteService.reposiroty.FindAllNotes()

	if len(notes) > 0 {
		for _, note := range notes {
			average += note.Nombre
		}
	}
	average /= len(notes)
	return entities.GlobalAverage{
		average,
		"Global average of promotion",
	}
}

func (noteService NoteService) ComputeAverageForEachUnites() []entities.UniteAverage {
	var uniteAverages []entities.UniteAverage

	unites := noteService.reposiroty.FindAllUnites()
	courses := noteService.reposiroty.FindAllCourses()

	for _, unite := range unites {
		var notes []entities.Note
		var courseAverage entities.UniteAverage
		notes = noteService.reposiroty.FindAllNotes()
		courseAverageResult := 0
		compteur := 0
		if len(notes) > 0 {
			for _, course := range courses {
				if course.UniteID == unite.ID {
					for _, note := range notes {
						if note.CourseName == course.Name {
							courseAverageResult += note.Nombre
							compteur++
						}
					}
				}
			}
			courseAverage.Average = courseAverageResult / compteur
			courseAverage.UniteName = unite.Name
			uniteAverages = append(uniteAverages, courseAverage)
		}
	}
	return uniteAverages
}

func (noteService NoteService) ComputeAverageForEachCourse() []entities.CourseAverage {
	var courseAverages []entities.CourseAverage

	courses := noteService.reposiroty.FindAllCourses()

	var notes []entities.Note
	var courseAverage entities.CourseAverage
	notes = noteService.reposiroty.FindAllNotes()
	if len(notes) > 0 {
		for _, course := range courses {
			courseAverageResult := 0
			compteur := 0
			for _, note := range notes {
				if note.CourseName == course.Name {
					courseAverageResult += note.Nombre
					compteur++
				}
			}
			courseAverage.Average = courseAverageResult / compteur
			courseAverage.CourseName = course.Name
			courseAverages = append(courseAverages, courseAverage)
		}
	}
	return courseAverages
}
