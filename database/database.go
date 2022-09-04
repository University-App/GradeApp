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
		&entities.Student{}, &entities.Note{}, &entities.Course{},
	}
	// TODO handle error
	database.AutoMigrate(models...)

	return database
}
