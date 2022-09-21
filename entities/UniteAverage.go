package entities

type UniteAverage struct {
	ID        uint `json:"ID" gorm:"PrimaryKey;unique;autoIncrement"`
	UniteName string
	Average   int
}
