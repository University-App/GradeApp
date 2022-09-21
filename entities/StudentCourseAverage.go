package entities

type StudentCourseAverage struct {
	ID             uint `json:"ID" gorm:"PrimaryKey;unique;autoIncrement"`
	CourseName     string
	StudentAverage StudentAverage `gorm:"embedded"`
}
