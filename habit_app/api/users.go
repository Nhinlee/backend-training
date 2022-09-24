package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	db "habits.com/habit/db/sqlc"
	"habits.com/habit/utils"
)

type createUserRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=6"`
}

type createUserResponse struct {
	UserID            int64     `json:"user_id"`
	FirstName         string    `json:"first_name"`
	LastName          string    `json:"last_name"`
	Email             string    `json:"email"`
	CreatedAt         time.Time `json:"created_at"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	arg := db.CreateUserParams{
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		Email:          req.Email,
		HashedPassword: hashedPassword,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		// Check violate pk
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code.Name() == "unique_violation" {
				ctx.JSON(http.StatusForbidden, errorResponse(pqErr))
				return
			}
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	userResponse := createUserResponse{
		UserID:            user.UserID,
		Email:             user.Email,
		FirstName:         user.FirstName,
		LastName:          user.LastName,
		PasswordChangedAt: user.PasswordChangedAt,
		CreatedAt:         user.CreatedAt,
	}

	ctx.JSON(http.StatusOK, userResponse)
}
