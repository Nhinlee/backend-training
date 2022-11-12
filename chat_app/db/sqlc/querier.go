// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"context"
)

type Querier interface {
	CreateConversation(ctx context.Context, arg CreateConversationParams) (Conversation, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteConversation(ctx context.Context, conversationID string) error
	DeleteUser(ctx context.Context, userID string) error
	GetConversationById(ctx context.Context, conversationID string) (Conversation, error)
	GetUser(ctx context.Context, email string) (User, error)
	ListUserIdByConversationId(ctx context.Context, conversationID string) ([]string, error)
	UpdateConversationInfo(ctx context.Context, arg UpdateConversationInfoParams) (Conversation, error)
	UpdateUserInfo(ctx context.Context, arg UpdateUserInfoParams) (User, error)
}

var _ Querier = (*Queries)(nil)
