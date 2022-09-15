package entities

type Unite struct {
	ID      uint     `json:"ID" gorm:"PrimaryKey;unique;autoIncrement"`
	Name    string   `json:"name"`
	Courses []Course `json:"courses"`
}
