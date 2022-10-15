package api

import (
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	grade_database "github.com/paulmarie/univesity/grade_app/persistence/sqlc"
)

type Server struct {
	store  *grade_database.Store
	router *gin.Engine
}

func NewServer(store *grade_database.Store) *Server {
	server := &Server{
		store: store,
	}
	router := gin.Default()

	router.GET("/coursesaverages", server.getCoursesAverages)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
