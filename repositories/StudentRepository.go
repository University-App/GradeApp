package repositories

import (
	"fmt"
	"github.com/paulmarie/univesity/grade_app/entities"
	"gorm.io/gorm"
)

type StudentRepository struct {
	DB *gorm.DB
}

func NewStudentRepository(db *gorm.DB) StudentRepository {
	return StudentRepository{db}
}

func (studentRepository StudentRepository) FindAllStudents() []entities.Student {
	var students []entities.Student

	if result := studentRepository.DB.Find(&students); result.Error != nil {
		fmt.Println(result.Error)
	}
	return students
}

func (studentRepository StudentRepository) AddStudent(student *entities.Student) entities.Student {

	if result := studentRepository.DB.Create(student); result.Error != nil {
		fmt.Println(result.Error)
	}
	return *student
}
