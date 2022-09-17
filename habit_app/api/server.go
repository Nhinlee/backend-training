package api

import (
	"github.com/gin-gonic/gin"
	db "habits.com/habit/db/sqlc"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/skills/create", server.createSkill)
	router.GET("/skills", server.getSkillsByUser)

	server.router = router
	return server
}

// Run http server on address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
