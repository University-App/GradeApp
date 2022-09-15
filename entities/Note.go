package entities

type Note struct {
	ID         uint   `json:"ID" gorm:"PrimaryKey;unique;autoIncrement"`
	Nombre     int    `json:"nombre"`
	StudentID  uint   `json:"studentID"`
	CourseName string `json:"courseName"`
}
