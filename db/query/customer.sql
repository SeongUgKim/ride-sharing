-- name: CreateCustomer :one
INSERT INTO customers (
    id,
    username,
    hashed_password,
    full_name,
    email
) VALUES (
     $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetCustomer :one
SELECT * FROM customers
WHERE id = $1 LIMIT 1;

-- name: DeleteCustomer :exec
DELETE FROM customers WHERE id = $1;