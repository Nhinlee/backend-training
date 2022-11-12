-- name: CreateUser :one
INSERT INTO users (
    first_name,
    last_name,
    email,
    hashed_password
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: UpdateUserInfo :one
UPDATE users SET first_name = $2, last_name = $3
WHERE user_id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users WHERE user_id = $1;