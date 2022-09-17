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

func (noteController NoteController) AddNote(ctx *fiber.Ctx) error {
	note := new(entities.Note)

	if err := ctx.BodyParser(note); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return ctx.JSON(noteController.noteServices.AddNote(note))
}

func (noteController NoteController) GetStudentsGlobalAverages(ctx *fiber.Ctx) error {
	return ctx.JSON(noteController.noteServices.ComputeGlobalAverageForEachStudent())
}

func (noteController NoteController) GetStudentsUnitesAverages(ctx *fiber.Ctx) error {
	return ctx.JSON(noteController.noteServices.ComputeCourseAverageForEachStudent())
}

func (noteController NoteController) GetStudentsCoursesAverages(ctx *fiber.Ctx) error {
	return ctx.JSON(noteController.noteServices.ComputeCourseAverageForEachStudent())
}

func (noteController NoteController) GetGobalAverage(ctx *fiber.Ctx) error {
	return ctx.JSON(noteController.noteServices.ComputeGlobalAverage())
}

func (noteController NoteController) GetUniteAverages(ctx *fiber.Ctx) error {
	return ctx.JSON(noteController.noteServices.ComputeAverageForEachUnites())
}

func (noteController NoteController) GetCoursesAverages(ctx *fiber.Ctx) error {
	return ctx.JSON(noteController.noteServices.ComputeAverageForEachCourse())
}
