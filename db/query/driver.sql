-- name: CreateDriver :one
INSERT INTO drivers (
    id,
    username,
    hashed_password,
    full_name,
    email,
    cab_id,
    dob
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
) RETURNING *;

-- name: GetDriver :one
SELECT * FROM drivers
WHERE id = $1 LIMIT 1;

-- name: DeleteDriver :exec
DELETE FROM drivers WHERE username = $1;

-- name: UpdateDriver :one
UPDATE drivers
SET cab_id = $2
WHERE id = $1
RETURNING *;