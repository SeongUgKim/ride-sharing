-- name: CreateTrip :one
INSERT INTO trips (
  id,
  customer_id,
  driver_id,
  status
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: GetTrip :one
SELECT * FROM trips
WHERE customer_id = $1 AND driver_id = $2
ORDER BY created_at DESC
LIMIT 1;

-- name: UpdateTrip :one
UPDATE trips
SET status = $4
WHERE customer_id = $1 AND driver_id = $2 AND status = $3
RETURNING *;

-- name: DeleteTrip :exec
DELETE FROM trips
WHERE customer_id = $1 AND driver_id = $2 AND status = $3;
