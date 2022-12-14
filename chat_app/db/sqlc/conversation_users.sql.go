// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: conversation_users.sql

package db

import (
	"context"
)

const createConversationUser = `-- name: CreateConversationUser :one
INSERT INTO conversation_users (
    user_id,
    conversation_id
) VALUES (
    $1, $2
) RETURNING user_id, conversation_id, status, created_at
`

type CreateConversationUserParams struct {
	UserID         string `json:"user_id"`
	ConversationID string `json:"conversation_id"`
}

func (q *Queries) CreateConversationUser(ctx context.Context, arg CreateConversationUserParams) (ConversationUser, error) {
	row := q.db.QueryRowContext(ctx, createConversationUser, arg.UserID, arg.ConversationID)
	var i ConversationUser
	err := row.Scan(
		&i.UserID,
		&i.ConversationID,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const listUserIdByConversationId = `-- name: ListUserIdByConversationId :many
SELECT user_id FROM conversation_users
WHERE conversation_id = $1
`

func (q *Queries) ListUserIdByConversationId(ctx context.Context, conversationID string) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, listUserIdByConversationId, conversationID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []string{}
	for rows.Next() {
		var user_id string
		if err := rows.Scan(&user_id); err != nil {
			return nil, err
		}
		items = append(items, user_id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
