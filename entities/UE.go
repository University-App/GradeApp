package entities

import (
	"github.com/google/uuid"
)

type Unite struct {
	ID      uuid.UUID `json:"ID" gorm:"PrimaryKey;unique"`
	Name    string    `json:"name"`
	Courses []Course  `json:"courses" gorm:"foreignKey:UniteID"`
}
