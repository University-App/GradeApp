package database

import (
	"github.com/google/uuid"
	"github.com/paulmarie/univesity/grade_app/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func Init() *gorm.DB {

	databaseURL := "postgres://user:password@university.grade.database.fr/grade_database"

	database, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	models := []interface{}{
		&entities.Student{}, &entities.Unite{}, &entities.Course{}, &entities.Note{},
	}

	err1 := database.Migrator().AutoMigrate(models...)
	if err1 != nil {
		panic("Cannot migrate models to the database ... ")
	}

	seedDatabase(database)

	return database
}

func seedDatabase(db *gorm.DB) {

	students := []entities.Student{
		{
			ID:        uuid.New(),
			LastName:  "LN1",
			FirstName: "FN1",
		},
		{
			ID:        uuid.New(),
			LastName:  "LN2",
			FirstName: "FN2",
		},
	}

	for _, student := range students {
		db.Create(&student)
	}

	ues := []entities.Unite{
		{
			ID:   uuid.New(),
			Name: "UE1",
		},
		{
			ID:   uuid.New(),
			Name: "UE2",
		},
	}

	courses := []entities.Course{
		{
			ID:   uuid.New(),
			Name: "Course1",
		},
		{
			ID:   uuid.New(),
			Name: "Course2",
		},
		{
			ID:   uuid.New(),
			Name: "Course3",
		},
		{
			ID:   uuid.New(),
			Name: "Course4",
		},
	}

	for _, ue := range ues {
		db.Create(&ue)
	}
	for index := range courses {
		if index%2 == 0 {
			db.Model(&ues[0]).Association("Courses").Append(&courses[index])
		} else {
			db.Model(&ues[1]).Association("Courses").Append(&courses[index])
		}
	}

	notes := []entities.Note{
		{
			ID:        uuid.New(),
			Nombre:    19,
			StudentID: students[0].ID,
		},
		{
			ID:        uuid.New(),
			Nombre:    10,
			StudentID: students[0].ID,
		},
		{
			ID:        uuid.New(),
			Nombre:    15,
			StudentID: students[0].ID,
		},
		{
			ID:        uuid.New(),
			Nombre:    18,
			StudentID: students[0].ID,
		},
	}
	for index := range courses {
		db.Model(&courses[index]).Association("Notes").Append(&notes)
	}
	for index := range students {
		db.Model(&students[index]).Association("Courses").Append(&courses)
	}

	//
	//var course1 entities.Course
	//db.First(&course1, "Name = ?", "Course1")
	//db.Model(&students[0]).Association("Courses").Append(&course1)
	//for index := range students {
	//}
}
