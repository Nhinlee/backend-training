-- name: CreateHabit :one
INSERT INTO habits (
    user_id,
    skill_id,
    title,
    max_consecutive_days
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetHabitsByUserAndSkill :many
SELECT * FROM habits
WHERE user_id = $1 AND skill_id = $2;

-- name: GetHabitsByUser :many
SELECT * FROM habits
WHERE user_id = $1;

-- name: DeleteHabit :exec
DELETE FROM habits WHERE habit_id = $1;