package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/paulmarie/univesity/grade_app/entities"
	"github.com/paulmarie/univesity/grade_app/services"
	"gorm.io/gorm"
)

type NoteController struct {
	noteServices services.NoteService
}

func NewNoteController(db *gorm.DB) NoteController {
	return NoteController{services.NewNoteService(db)}
}

func (noteController NoteController) HelloWorld(ctx *fiber.Ctx) error {
	return ctx.SendString("Hello World!")
}

func (noteController NoteController) GetAllNotes(ctx *fiber.Ctx) error {

	return ctx.JSON(noteController.noteServices.GetAllNotes())
}

func (noteController NoteController) AddNote(ctx *fiber.Ctx) error {
	note := new(entities.Note)

	if err := ctx.BodyParser(note); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return ctx.JSON(noteController.noteServices.AddNote(note))
}

func (noteController NoteController) GetAverage(ctx *fiber.Ctx) error {
	return ctx.JSON(noteController.noteServices.ComputeAverage())
}
