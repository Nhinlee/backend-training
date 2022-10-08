package rpcapi

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	db "habits.com/habit/db/sqlc"
	"habits.com/habit/pb"
)

func convertUserToUserRes(user db.User) *pb.User {
	return &pb.User{
		FirstName:         user.FirstName,
		LastName:          user.LastName,
		Email:             user.Email,
		PasswordChangedAt: timestamppb.New(user.PasswordChangedAt),
		CreatedAt:         timestamppb.New(user.CreatedAt),
	}
}
