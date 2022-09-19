-- name: CreateHabitLog :one
INSERT INTO habit_logs (
    user_id,
    habit_id,
    date_time
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetHabitLogsByUser :many
SELECT * FROM habit_logs
WHERE user_id = $1
ORDER BY habit_id;

-- name: GetLatestHabitLogByUser :one
SELECT * from habit_logs
WHERE user_id = $1
ORDER BY date_time DESC
LIMIT 1
OFFSET 1;