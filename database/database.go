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
		&entities.GlobalAverage{}, &entities.CourseAverage{}, &entities.UniteAverage{},
		&entities.StudentUniteAverage{}, &entities.StudentCourseAverage{}, &entities.StudentGlobalAverage{},
		&entities.StudentAverage{},
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

	notesStudent1 := []entities.Note{
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
	notesStudent2 := []entities.Note{
		{
			Nombre:     15,
			CourseName: "Course1",
		},
		{
			Nombre:     8,
			CourseName: "Course2",
		},
		{
			Nombre:     10,
			CourseName: "Course3",
		},
		{
			Nombre:     20,
			CourseName: "Course4",
		},
	}

	var studentsFromDB []entities.Student
	db.Find(&studentsFromDB)
	for index := range studentsFromDB {
		db.Model(&studentsFromDB[index]).Association("Courses").Append(&coursesFromDB)
	}

	for indexNote := range notesStudent1 {
		db.Model(&studentsFromDB[0]).Association("Notes").Append([]entities.Note{notesStudent1[indexNote]})
	}

	for indexNote := range notesStudent2 {
		db.Model(&studentsFromDB[1]).Association("Notes").Append([]entities.Note{notesStudent2[indexNote]})
	}

}
