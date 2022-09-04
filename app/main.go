package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/paulmarie/univesity/grade_app/controllers"
	"github.com/paulmarie/univesity/grade_app/database"
)

func SetUpRoutes(app *fiber.App, controller controllers.NoteController) {
	app.Get("/", controller.HelloWorld)
	app.Get("/allNotes", controller.GetAllNotes)
	app.Post("/newNote", controller.AddNote)
	app.Get("/average", controller.GetAverage)
}

func main() {
	DB := database.Init()

	noteController := controllers.NewNoteController(DB)

	app := fiber.New()

	SetUpRoutes(app, noteController)

	err := app.Listen(":4000")
	if err != nil {
		panic("Cannot start the application")
	}
}
