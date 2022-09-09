package entities

import "github.com/google/uuid"

type Student struct {
	ID        uuid.UUID `json:"studentID" gorm:"PrimaryKey;unique"`
	LastName  string    `json:"lastName"`
	FirstName string    `json:"firstName"`
	Courses   []Course  `json:"courses" gorm:"many2many:students_courses;foreignKey:ID;Reference:ID;"`
}
