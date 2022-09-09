package entities

import (
	"github.com/google/uuid"
)

type Course struct {
	ID      uuid.UUID `json:"ID" gorm:"PrimaryKey;unique"`
	UniteID uuid.UUID `json:"ueid" gorm:"foreignKey:ID"`
	Name    string    `json:"name"`
	Notes   []Note    `json:"notes" gorm:"foreignKey:CourseID"`
}
