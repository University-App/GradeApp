package services

import (
	"github.com/paulmarie/univesity/grade_app/entities"
	"github.com/paulmarie/univesity/grade_app/repositories"
	"gorm.io/gorm"
	"sort"
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

func (noteService NoteService) ComputeGlobalAverageForEachStudent() []entities.StudentGlobalAverage {
	var studentsGlobalAverages []entities.StudentGlobalAverage

	students := noteService.reposiroty.FindAllStudents()

	for _, student := range students {
		gobalAverageResult := 0
		var notes []entities.Note
		var globalAverage entities.StudentGlobalAverage
		notes = noteService.reposiroty.FindAllNotesOfStudent(student.ID)
		if len(notes) > 0 {
			for _, note := range notes {
				gobalAverageResult += note.Nombre
			}

			globalAverage.StudentAverage.Average = gobalAverageResult / len(notes)
		}
		globalAverage.StudentAverage.StudentName = student.FirstName + "-" + student.LastName

		studentsGlobalAverages = append(studentsGlobalAverages, globalAverage)
		noteService.reposiroty.CreateStudentGlobalAverage(globalAverage)
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
			var studentUniteAverage entities.StudentUniteAverage
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
				studentUniteAverage.StudentAverage.Average = courseAverageResult / compteur
				studentUniteAverage.UniteName = unite.Name
				studentUniteAverage.StudentAverage.StudentName = student.FirstName + "-" + student.LastName
				studentUniteAverages = append(studentUniteAverages, studentUniteAverage)
				noteService.reposiroty.CreateStudentUniteAverage(studentUniteAverage)
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
		var studentCourseAverage entities.StudentCourseAverage
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
				studentCourseAverage.StudentAverage.Average = courseAverageResult / compteur
				studentCourseAverage.StudentAverage.StudentName = student.FirstName + "-" + student.LastName
				studentCourseAverage.CourseName = course.Name
				studentCourseAverages = append(studentCourseAverages, studentCourseAverage)
				noteService.reposiroty.CreateStudentCourseAverage(studentCourseAverage)
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
		Average:       average,
		PromotionName: "Global average of promotion",
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

func (noteService NoteService) DetermineGlobalRankForEachStudent() []entities.GlobalRank {
	studentGlobalAverages := noteService.reposiroty.FindAllStudentGlobalAverage()
	var globalRanks []entities.GlobalRank
	sort.Slice(studentGlobalAverages, func(i, j int) bool {
		return studentGlobalAverages[i].StudentAverage.Average > studentGlobalAverages[j].StudentAverage.Average
	})

	for index := range studentGlobalAverages {
		var globalRank entities.GlobalRank

		globalRank.Rank = index + 1
		globalRank.StudentAverage = studentGlobalAverages[index].StudentAverage
		globalRanks = append(globalRanks, globalRank)
	}
	return globalRanks
}

func (noteService NoteService) DetermineUnitesRankForEachStudent() [][]entities.GlobalRank {
	studentGlobalAverages := noteService.reposiroty.FindAllStudentUnitesAverage()
	unites := noteService.reposiroty.FindAllUnites()

	var globalRankss [][]entities.GlobalRank
	for _, unite := range unites {
		var averages []entities.StudentUniteAverage
		var globalRanks []entities.GlobalRank
		for _, studentGlobalAverage := range studentGlobalAverages {
			if studentGlobalAverage.UniteName == unite.Name {
				averages = append(averages, studentGlobalAverage)
			}
		}
		sort.Slice(averages, func(i, j int) bool {
			return averages[i].StudentAverage.Average > averages[j].StudentAverage.Average
		})
		for index := range averages {
			var globalRank entities.GlobalRank

			globalRank.Rank = index + 1
			globalRank.StudentAverage = averages[index].StudentAverage
			globalRanks = append(globalRanks, globalRank)
		}
		globalRankss = append(globalRankss, globalRanks)
	}
	return globalRankss
}

func (noteService NoteService) DetermineCoursesRankForEachStudent() [][]entities.GlobalRank {
	studentCourseAverages := noteService.reposiroty.FindAllStudentCoursesAverage()
	courses := noteService.reposiroty.FindAllCourses()

	var globalRankss [][]entities.GlobalRank
	for _, course := range courses {
		var averages []entities.StudentCourseAverage
		var globalRanks []entities.GlobalRank
		for _, studentCourseAverage := range studentCourseAverages {
			if studentCourseAverage.CourseName == course.Name {
				averages = append(averages, studentCourseAverage)
			}
		}
		sort.Slice(averages, func(i, j int) bool {
			return averages[i].StudentAverage.Average > averages[j].StudentAverage.Average
		})
		for index := range averages {
			var globalRank entities.GlobalRank

			globalRank.Rank = index + 1
			globalRank.StudentAverage = averages[index].StudentAverage
			globalRanks = append(globalRanks, globalRank)
		}
		globalRankss = append(globalRankss, globalRanks)
	}
	return globalRankss
}
