package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "habits.com/habit/db/sqlc"
)

type createSkillRequest struct {
	UserID int64  `json:"user_id" binding:"required"`
	Title  string `json:"title" binding:"required"`
}

func (server *Server) createSkill(ctx *gin.Context) {
	var req createSkillRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateSkillParams{
		UserID: req.UserID,
		Title:  req.Title,
	}

	skill, err := server.store.CreateSkill(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, skill)
}
