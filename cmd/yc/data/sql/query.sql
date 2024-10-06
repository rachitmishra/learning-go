-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY created_at;

-- name: ListLinks :many
SELECT * FROM links
ORDER BY id;

-- name: CreateUser :one
INSERT INTO users (
  username, email, password_hash
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: CreateLink :one
INSERT INTO links (
  user_id, original_url, short_url, clicks
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: DeleteLink :exec
DELETE FROM links
WHERE id = $1;