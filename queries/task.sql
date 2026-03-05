-- name: CreateTask :one
INSERT INTO tasks (user_id, name, description, date, time_start, time_end)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetTaskByID :one
SELECT * FROM tasks
WHERE id = $1 AND user_id = $2 AND deleted_at IS NULL;

-- name: UpdateTask :one
UPDATE tasks
SET name = $3, description = $4, date = $5, time_start = $6, time_end = $7, updated_at = NOW()
WHERE id = $1 AND user_id = $2 AND deleted_at IS NULL
RETURNING *;

-- name: SetTaskCompleted :one
UPDATE tasks SET is_completed = $3, updated_at = NOW()
WHERE id = $1 AND user_id = $2
RETURNING *;

-- name: SetTaskFavourite :one
UPDATE tasks SET is_favourite = $3, updated_at = NOW()
WHERE id = $1 AND user_id = $2
RETURNING *;

-- name: SoftDeleteTask :exec
UPDATE tasks SET deleted_at = NOW()
WHERE id = $1 AND user_id = $2;

-- name: RestoreTask :exec
UPDATE tasks SET deleted_at = NULL
WHERE id = $1 AND user_id = $2;

-- name: AddTagToTask :exec
INSERT INTO task_tags (task_id, tag_id) VALUES ($1, $2)
ON CONFLICT DO NOTHING;

-- name: RemoveTagFromTask :exec
DELETE FROM task_tags WHERE task_id = $1 AND tag_id = $2;

-- name: GetTagsByTaskID :many
SELECT t.* FROM tags t
JOIN task_tags tt ON tt.tag_id = t.id
WHERE tt.task_id = $1;