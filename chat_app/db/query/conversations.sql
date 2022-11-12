-- name: CreateConversation :one
INSERT INTO conversations (
    conversation_id,
    conversation_name
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetConversationById :one
SELECT * FROM conversations
WHERE conversation_id = $1 LIMIT 1;


-- name: UpdateConversationInfo :one
UPDATE conversations SET conversation_name = $2
WHERE conversation_id = $1
RETURNING *;

-- name: DeleteConversation :exec
DELETE FROM conversations WHERE conversation_id = $1;