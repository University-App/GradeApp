package database

import (
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
			LastName:  "LN1",
			FirstName: "FN1",
		},
		{
			LastName:  "LN2",
			FirstName: "FN2",
		},
	}

	for _, student := range students {
		db.Create(&student)
	}

	ues := []entities.Unite{
		{
			Name: "UE1",
		},
		{
			Name: "UE2",
		},
	}

	courses := []entities.Course{
		{
			Name: "Course1",
		},
		{
			Name: "Course2",
		},
		{
			Name: "Course3",
		},
		{
			Name: "Course4",
		},
	}

	for _, ue := range ues {
		db.Create(&ue)
	}
	var uesFromDB []entities.Unite
	db.Find(&uesFromDB)

	for index := range courses {
		if index%2 == 0 {
			db.Model(&uesFromDB[0]).Association("Courses").Append(&courses[index])
		} else {
			db.Model(&uesFromDB[1]).Association("Courses").Append(&courses[index])
		}
	}

	var coursesFromDB []entities.Course
	db.Find(&coursesFromDB)

	notes := []entities.Note{
		{
			Nombre:     19,
			CourseName: "Course1",
		},
		{
			Nombre:     10,
			CourseName: "Course2",
		},
		{
			Nombre:     15,
			CourseName: "Course3",
		},
		{
			Nombre:     18,
			CourseName: "Course4",
		},
	}

	var studentsFromDB []entities.Student
	db.Find(&studentsFromDB)
	for index := range studentsFromDB {
		db.Model(&studentsFromDB[index]).Association("Courses").Append(&coursesFromDB)
	}
	for index := range studentsFromDB {
		for indexNote := range notes {
			db.Model(&studentsFromDB[index]).Association("Notes").Append([]entities.Note{notes[indexNote]})
		}
	}
}
