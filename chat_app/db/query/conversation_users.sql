-- name: CreateConversationUser :one
INSERT INTO conversation_users (
    user_id,
    conversation_id
) VALUES (
    $1, $2
) RETURNING *;

-- name: ListUserIdByConversationId :many
SELECT user_id FROM conversation_users
WHERE conversation_id = $1;
