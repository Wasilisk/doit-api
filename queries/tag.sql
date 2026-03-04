-- name: CreateTag :one
INSERT INTO tags (user_id, name, color)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetTagByID :one
SELECT * FROM tags
WHERE id = $1 AND user_id = $2;

-- name: GetTagsByUserID :many
SELECT * FROM tags
WHERE user_id = $1
ORDER BY created_at DESC;

-- name: UpdateTag :one
UPDATE tags
SET name = $3, color = $4, updated_at = NOW()
WHERE id = $1 AND user_id = $2
RETURNING *;

-- name: DeleteTag :exec
DELETE FROM tags
WHERE id = $1 AND user_id = $2;