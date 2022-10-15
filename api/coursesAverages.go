package api

import (
	"github.com/gin-gonic/gin"
	grade_database "github.com/paulmarie/univesity/grade_app/persistence/sqlc"
	"net/http"
)

type getCourseAverageRequest struct {
	CourseName string `form:"courseName" binding:"required"`
}

func (server *Server) getCoursesAverages(ctx *gin.Context) {

	var req getCourseAverageRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	arg := grade_database.CourseAverageTxParams{
		CourseName: req.CourseName,
	}

	courseAverage, errGet := server.store.CourseAverageTx(ctx, arg)
	if errGet != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(errGet))
	}

	ctx.JSON(http.StatusOK, courseAverage)
}
