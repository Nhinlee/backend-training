-- name: ListUserIdByConversationId :many
SELECT user_id FROM conversation_members
WHERE conversation_id = $1;
