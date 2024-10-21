-- name: CreateURL :exec
INSERT INTO url (original, shorten)
VALUES (?, ?);

-- name: GetOriginalURL :one
SELECT 
    original
FROM url
WHERE shorten = ? LIMIT 1;
