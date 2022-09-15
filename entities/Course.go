package entities

type Course struct {
	ID      uint   `json:"ID" gorm:"PrimaryKey;unique;autoIncrement"`
	UniteID uint   `json:"ueid"`
	Name    string `json:"name"`
}
