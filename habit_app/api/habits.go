package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "habits.com/habit/db/sqlc"
)

type CreateHabitRequest struct {
	UserID                int64  `json:"user_id" binding:"required"`
	SkillID               int64  `json:"skill_id"`
	Title                 string `json:"title" binding:"required"`
	TargetConsecutiveDays int32  `json:"target_consecutive_days" binding:"required"`
}

func (server *Server) createHabit(ctx *gin.Context) {
	var req CreateHabitRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	habit, err := server.store.CreateHabit(ctx, db.CreateHabitParams{
		UserID: req.UserID,
		SkillID: sql.NullInt64{
			Int64: req.SkillID,
			Valid: req.SkillID != 0,
		},
		Title:                 req.Title,
		TargetConsecutiveDays: req.TargetConsecutiveDays,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, habit)
}
