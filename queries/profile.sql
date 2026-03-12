-- name: CreateProfile :one
INSERT INTO user_profiles (user_id, full_name, avatar_url)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetProfileByUserID :one
SELECT p.*, u.email FROM user_profiles p
JOIN users u ON u.id = p.user_id
WHERE p.user_id = $1;

-- name: UpdateProfile :one
UPDATE user_profiles
SET
    full_name  = COALESCE(sqlc.narg(full_name), full_name),
    avatar_url = COALESCE(sqlc.narg(avatar_url), avatar_url),
    updated_at = NOW()
WHERE user_id = sqlc.arg(user_id)
RETURNING *;