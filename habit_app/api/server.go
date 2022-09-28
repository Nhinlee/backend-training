package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	db "habits.com/habit/db/sqlc"
	"habits.com/habit/token"
	"habits.com/habit/utils"
)

type Server struct {
	config       *utils.Config
	store        db.Store
	router       *gin.Engine
	tokenFactory token.TokenFactory
}

func NewServer(config *utils.Config, store db.Store) (*Server, error) {
	tokenFactory, err := token.NewPasetoTokenFactory(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token factory: %w", err)
	}

	server := &Server{
		store:        store,
		tokenFactory: tokenFactory,
		config:       config,
	}

	server.SetupRouter()

	return server, nil
}

func (server *Server) SetupRouter() {
	router := gin.Default()

	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)

	router.POST("/skills", server.createSkill)
	router.GET("/skills", server.getSkillsByUser)

	router.POST("/habits", server.createHabit)
	router.POST("/habit_logs", server.createHabitLog)

	server.router = router
}

// Run http server on address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
