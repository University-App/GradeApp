package entities

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	Name    string ` json:"name"`
	Notes   []Note `json:"notes"`
	UniteID uint   `json:"ueid"`
}
