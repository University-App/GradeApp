package entities

type CourseAverage struct {
	ID         uint `json:"ID" gorm:"PrimaryKey;unique;autoIncrement"`
	CourseName string
	Average    int
}
