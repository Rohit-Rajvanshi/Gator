-- name: GetFeedID :one

SELECT id FROM feeds WHERE feeds.url = $1 LIMIT 1;