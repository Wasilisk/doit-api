-- name: CreateTag :one
INSERT INTO tags (user_id, name, color)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetTagByID :one
SELECT * FROM tags
WHERE id = $1 AND user_id = $2;

-- name: GetTagsByUserID :many
SELECT t.id, t.user_id, t.name, t.color, t.created_at, t.updated_at,
    COUNT(tt.task_id)::int AS task_count
FROM tags t
LEFT JOIN task_tags tt ON t.id = tt.tag_id
WHERE t.user_id = $1
GROUP BY t.id
ORDER BY t.created_at DESC;

-- name: UpdateTag :one
UPDATE tags
SET name = $3, color = $4, updated_at = NOW()
WHERE id = $1 AND user_id = $2
RETURNING *;

-- name: DeleteTag :exec
DELETE FROM tags
WHERE id = $1 AND user_id = $2;