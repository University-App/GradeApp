package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/paulmarie/univesity/grade_app/entities"
	"github.com/paulmarie/univesity/grade_app/services"
	"gorm.io/gorm"
)

type StudentController struct {
	studentService services.StudentService
}

func NewStudentController(db *gorm.DB) StudentController {
	return StudentController{services.NewStudentService(db)}
}

func (noteController StudentController) GetAllStudents(ctx *fiber.Ctx) error {

	return ctx.JSON(noteController.studentService.GetAllStudents())
}

func (noteController StudentController) AddStudent(ctx *fiber.Ctx) error {
	student := new(entities.Student)

	if err := ctx.BodyParser(student); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return ctx.JSON(noteController.studentService.AddStudent(student))
}
