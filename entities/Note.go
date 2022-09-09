package entities

import (
	"github.com/google/uuid"
)

type Note struct {
	ID        uuid.UUID `json:"ID" gorm:"PrimaryKey;unique"`
	Nombre    int       `json:"nombre"`
	StudentID uuid.UUID `json:"studentId" gorm:"foreignKey:ID"`
	CourseID  uuid.UUID `json:"courseId" gorm:"foreignKey:ID"`
}
