-- name: CreateRating :one
INSERT INTO ratings (
 id,
 customer_id,
 driver_id,
 trip_id,
 rating,
 feedback
) VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetRating :one
SELECT * FROM ratings
WHERE customer_id = $1 AND driver_id = $2 AND trip_id = $3;

-- name: DeleteRating :exec
DELETE FROM ratings
WHERE customer_id = $1 AND driver_id = $2 AND trip_id = $3;