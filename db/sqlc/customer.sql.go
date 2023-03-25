// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: customer.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createCustomer = `-- name: CreateCustomer :one
INSERT INTO customers (
    id,
    username,
    hashed_password,
    full_name,
    email
) VALUES (
     $1, $2, $3, $4, $5
) RETURNING id, username, hashed_password, full_name, email, password_changed_at, created_at
`

type CreateCustomerParams struct {
	ID             uuid.UUID `json:"id"`
	Username       string    `json:"username"`
	HashedPassword string    `json:"hashed_password"`
	FullName       string    `json:"full_name"`
	Email          string    `json:"email"`
}

func (q *Queries) CreateCustomer(ctx context.Context, arg CreateCustomerParams) (Customer, error) {
	row := q.db.QueryRowContext(ctx, createCustomer,
		arg.ID,
		arg.Username,
		arg.HashedPassword,
		arg.FullName,
		arg.Email,
	)
	var i Customer
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.PasswordChangedAt,
		&i.CreatedAt,
	)
	return i, err
}

const deleteCustomer = `-- name: DeleteCustomer :exec
DELETE FROM customers WHERE id = $1
`

func (q *Queries) DeleteCustomer(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteCustomer, id)
	return err
}

const getCustomer = `-- name: GetCustomer :one
SELECT id, username, hashed_password, full_name, email, password_changed_at, created_at FROM customers
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetCustomer(ctx context.Context, id uuid.UUID) (Customer, error) {
	row := q.db.QueryRowContext(ctx, getCustomer, id)
	var i Customer
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.PasswordChangedAt,
		&i.CreatedAt,
	)
	return i, err
}
