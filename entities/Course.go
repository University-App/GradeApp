package entities

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	Name     string    ` json:"name"`
	Notes    []Note    `json:"notes"`
	Students []Student `json:"students" gorm:"many2many: student_courses"`
}
