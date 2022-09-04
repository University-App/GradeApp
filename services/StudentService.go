package services

import (
	"github.com/paulmarie/univesity/grade_app/entities"
	"github.com/paulmarie/univesity/grade_app/repositories"
	"gorm.io/gorm"
)

type StudentService struct {
	reposiroty repositories.StudentRepository
}

func NewStudentService(db *gorm.DB) StudentService {
	return StudentService{repositories.NewStudentRepository(db)}
}

func (studentService StudentService) GetAllStudents() []entities.Student {
	return studentService.reposiroty.FindAllStudents()
}

func (studentService StudentService) AddStudent(student *entities.Student) entities.Student {

	return studentService.reposiroty.AddStudent(student)
}
