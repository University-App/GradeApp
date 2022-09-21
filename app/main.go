package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/paulmarie/univesity/grade_app/controllers"
	"github.com/paulmarie/univesity/grade_app/database"
)

func SetUpRoutes(app *fiber.App, noteController controllers.NoteController, studentController controllers.StudentController) {
	app.Post("/newNote", noteController.AddNote)
	app.Get("/allStudents", studentController.GetAllStudents)
	app.Post("/newStudent", studentController.AddStudent)

	app.Get("/studentsGlobalAverages", noteController.GetStudentsGlobalAverages)
	app.Get("/studentsUnitesAverages", noteController.GetStudentsUnitesAverages)
	app.Get("/studentsCoursesAverages", noteController.GetStudentsCoursesAverages)

	app.Get("/gobalAverage", noteController.GetGobalAverage)
	app.Get("/uniteAverages", noteController.GetUniteAverages)
	app.Get("/coursesAverages", noteController.GetCoursesAverages)

	app.Get("/globalRanks", noteController.GetGlobalRankStudents)
	app.Get("/unitesRanks", noteController.GetUnitesRankStudents)
	app.Get("/coursesRanks", noteController.GetCoursesRankStudents)
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
