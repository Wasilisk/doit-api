-- name: CreateUser :one
INSERT INTO users (email, password)
VALUES ($1, $2)
RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1;

-- name: GetUserWithProfile :one
SELECT 
    u.id,
    u.email,
    p.full_name,
    p.avatar_url,
    u.created_at
FROM users u
LEFT JOIN user_profiles p ON p.user_id = u.id
WHERE u.id = $1;