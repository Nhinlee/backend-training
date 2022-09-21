package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "habits.com/habit/db/sqlc"
)

type CreateHabitLogRequest struct {
	UserID  int64 `json:"user_id" binding:"required"`
	HabitID int64 `json:"habit_id" binding:"required"`
}

func (server *Server) createHabitLog(ctx *gin.Context) {
	var req CreateHabitLogRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	rs, err := server.store.CreateHabitLogTx(ctx, db.CreateHabitLogTxParams{
		UserID:  req.UserID,
		HabitID: req.HabitID,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, rs)
}
