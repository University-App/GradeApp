package entities

type Student struct {
	ID        uint     `json:"studentID" gorm:"PrimaryKey;unique;autoIncrement"`
	LastName  string   `json:"lastName"`
	FirstName string   `json:"firstName"`
	Notes     []Note   `json:"notes"`
	Courses   []Course `json:"courses" gorm:"many2many:students_courses;foreignKey:ID;Reference:ID;"`
}
