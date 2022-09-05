package entities

import "gorm.io/gorm"

type Unite struct {
	gorm.Model
	Name    string   `json:"name"`
	Courses []Course `json:"courses"`
}
