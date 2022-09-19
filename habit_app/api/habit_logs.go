package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateHabitLogRequest struct {
	UserID  int64 `json:"user_id"`
	HabitID int64 `json:"habit_id"`
}

func (server *Server) createHabitLog(ctx *gin.Context) {
	var req CreateHabitLogRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

}
