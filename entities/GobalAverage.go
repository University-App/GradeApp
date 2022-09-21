package entities

type GlobalAverage struct {
	ID            uint `json:"ID" gorm:"PrimaryKey;unique;autoIncrement"`
	Average       int
	PromotionName string
}
