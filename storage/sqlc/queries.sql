-- name: CreateURL :exec
INSERT INTO url (original, shorten_id)
VALUES (?, ?);

-- name: GetOriginalURL :one
SELECT 
    original
FROM url
WHERE shorten_id = ? LIMIT 1;
