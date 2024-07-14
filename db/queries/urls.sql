-- name: CreateURL :one
INSERT INTO urls (original_url, short_url) 
VALUES ($1, $2)
RETURNING id, original_url, short_url, created_at;

-- name: GetURLByShortURL :one
SELECT id, original_url, short_url, created_at 
FROM urls 
WHERE short_url = $1;


-- name: GetURLByOriginalURL :one
SELECT id, original_url, short_url, created_at
FROM urls
WHERE original_url = $1;

-- name: GetURLByID :one
SELECT id, original_url, short_url, created_at
FROM urls
WHERE id = $1;

-- name: GetURLs :many
SELECT id, original_url, short_url, created_at
FROM urls;

-- name: DeleteURL :exec
DELETE FROM urls
WHERE id = $1;



-- name: UpdateURL :one
UPDATE urls
SET original_url = $2, short_url = $3
WHERE id = $1
RETURNING id, original_url, short_url, created_at;


-- name: GetURLsByDate :many
SELECT id, original_url, short_url, created_at
FROM urls
WHERE created_at >= $1 AND created_at <= $2;

