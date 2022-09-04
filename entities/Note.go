package entities

type Note struct {
	ID     int `json:"id" gorm:"primaryKey;autoIncrement; unique; not null"`
	Nombre int `json:"nombre"`
}
