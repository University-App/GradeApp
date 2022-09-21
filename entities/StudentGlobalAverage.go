package entities

type StudentGlobalAverage struct {
	ID             uint           `json:"ID" gorm:"PrimaryKey;unique;autoIncrement"`
	StudentAverage StudentAverage `gorm:"embedded"`
}
