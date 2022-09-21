package entities

type StudentUniteAverage struct {
	ID             uint `json:"ID" gorm:"PrimaryKey;unique;autoIncrement"`
	UniteName      string
	StudentAverage StudentAverage `gorm:"embedded"`
}
