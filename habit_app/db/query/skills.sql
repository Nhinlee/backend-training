-- name: CreateSkill :one
INSERT INTO skills (
    user_id,
    title
) VALUES (
    $1, $2
) RETURNING *;

-- name: ListSkillsByUser :many
SELECT * FROM skills
ORDER BY skill_id
LIMIT $1
OFFSET $2;

-- name: UpdateSkill :one
UPDATE skills SET title = $2
WHERE skill_id = $1
RETURNING *;

-- name: DeleteSkill :exec
DELETE FROM skills WHERE skill_id = $1;
