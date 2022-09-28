package api

import (
	"github.com/gin-gonic/gin"
	db "habits.com/habit/db/sqlc"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/users", server.createUser)

	router.POST("/skills", server.createSkill)
	router.GET("/skills", server.getSkillsByUser)

	router.POST("/habits", server.createHabit)
	router.POST("/habit_logs", server.createHabitLog)

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
