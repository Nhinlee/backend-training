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

type getSkillsByUserRequest struct {
	UserID   int64 `form:"user_id"`
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=20"`
}

func (server *Server) getSkillsByUser(ctx *gin.Context) {
	var req getSkillsByUserRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListSkillsByUserParams{
		UserID: req.UserID,
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	skills, err := server.store.ListSkillsByUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, skills)
}
