-- name: CreateCabs :one
INSERT INTO cabs (
  id,
  cab_type,
  reg_no
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetCab :one
SELECT * FROM cabs
WHERE id = $1 LIMIT 1;

-- name: DeleteCab :exec
DELETE FROM cabs WHERE id = $1;

-- name: UpdateCab :one
UPDATE cabs
SET cab_type = $2
WHERE id = $1
RETURNING *;