package db

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"uber/util"
)

func createRandomCustomer(t *testing.T) Customer {
	randomUUID, err := uuid.NewRandom()
	require.NoError(t, err)

	username := util.RandomName()

	hashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)

	arg := CreateCustomerParams{
		ID:             randomUUID,
		Username:       username,
		HashedPassword: hashedPassword,
		FullName:       username,
		Email:          util.RandomEmail(),
	}

	customer, err := testQueries.CreateCustomer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, customer)

	require.Equal(t, arg.ID, customer.ID)
	require.Equal(t, arg.Username, customer.Username)
	require.Equal(t, arg.HashedPassword, customer.HashedPassword)
	require.Equal(t, arg.FullName, customer.FullName)
	require.Equal(t, arg.Email, customer.Email)
	require.NotZero(t, customer.CreatedAt)
	require.True(t, customer.PasswordChangedAt.IsZero())

	return customer
}

func TestCreateCustomer(t *testing.T) {
	createRandomCustomer(t)
}

func TestGetCustomer(t *testing.T) {
	customer1 := createRandomCustomer(t)
	customer2, err := testQueries.GetCustomer(context.Background(), customer1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, customer2)

	require.Equal(t, customer1.ID, customer2.ID)
	require.Equal(t, customer1.Username, customer2.Username)
	require.Equal(t, customer1.HashedPassword, customer2.HashedPassword)
	require.Equal(t, customer1.FullName, customer2.FullName)
	require.Equal(t, customer1.Email, customer2.Email)
	require.WithinDuration(t, customer1.PasswordChangedAt, customer2.PasswordChangedAt, time.Second)
	require.WithinDuration(t, customer1.CreatedAt, customer2.CreatedAt, time.Second)
}
