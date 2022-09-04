package entities

import "gorm.io/gorm"

type Note struct {
	gorm.Model
	Nombre    int  `json:"nombre"`
	StudentID uint `json:"studentId"`
	CourseID  uint `json:"courseId"`
}
