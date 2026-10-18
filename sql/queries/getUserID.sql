-- name: GetUserID :one

SELECT id FROM users WHERE users.name = $1 LIMIT 1;

