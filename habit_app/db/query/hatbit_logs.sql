-- name: CreateHabitLog :one
INSERT INTO habit_logs (
    user_id,
    habit_id,
    created_at
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetHabitLogsByUser :many
SELECT * FROM habit_logs
WHERE user_id = $1
ORDER BY habit_id;

-- name: GetLatestHabitLogByUser :many
SELECT * from habit_logs
WHERE user_id = $1 AND habit_id = $2
ORDER BY created_at DESC
LIMIT 1;