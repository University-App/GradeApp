package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/paulmarie/univesity/grade_app/controllers"
	"github.com/paulmarie/univesity/grade_app/database"
)

func SetUpRoutes(app *fiber.App, noteController controllers.NoteController, studentController controllers.StudentController) {
	app.Get("/", noteController.HelloWorld)
	app.Get("/allNotes", noteController.GetAllNotes)
	app.Post("/newNote", noteController.AddNote)
	app.Get("/average", noteController.GetAverage)
	app.Get("/allStudents", studentController.GetAllStudents)
	app.Post("/newStudent", studentController.AddStudent)
}

func main() {
	DB := database.Init()

	noteController := controllers.NewNoteController(DB)
	studentControler := controllers.NewStudentController(DB)

	app := fiber.New()

	SetUpRoutes(app, noteController, studentControler)

	err := app.Listen(":4000")
	if err != nil {
		panic("Cannot start the application")
	}
}
