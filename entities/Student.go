package entities

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	LastName  string   `json:"lastName"`
	FirstName string   `json:"firstName"`
	Courses   []Course `json:"courses" gorm:"many2many:student_courses"`
}
