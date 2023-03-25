package db

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"uber/util"
)

func createRandomDriver(t *testing.T) Driver {
	randomID, err := uuid.NewRandom()
	require.NoError(t, err)

	randomCab := createRandomCab(t)

	hashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)

	username := util.RandomName()
	arg := CreateDriverParams{
		ID:             randomID,
		Username:       username,
		HashedPassword: hashedPassword,
		FullName:       username,
		Email:          util.RandomEmail(),
		CabID:          randomCab.ID,
		Dob:            time.Now(),
	}

	driver, err := testQueries.CreateDriver(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, driver)

	require.Equal(t, arg.ID, driver.ID)
	require.Equal(t, arg.Username, driver.Username)
	require.Equal(t, arg.HashedPassword, driver.HashedPassword)
	require.Equal(t, arg.FullName, driver.FullName)
	require.Equal(t, arg.Email, driver.Email)
	require.Equal(t, arg.CabID, driver.CabID)
	require.NotZero(t, driver.Dob)
	require.NotZero(t, driver.CreatedAt)
	require.True(t, driver.PasswordChangedAt.IsZero())

	return driver
}

func TestCreateDriver(t *testing.T) {
	createRandomDriver(t)
}

func TestGetDriver(t *testing.T) {
	driver1 := createRandomDriver(t)
	driver2, err := testQueries.GetDriver(context.Background(), driver1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, driver2)

	require.Equal(t, driver1.ID, driver2.ID)
	require.Equal(t, driver1.Username, driver2.Username)
	require.Equal(t, driver1.HashedPassword, driver2.HashedPassword)
	require.Equal(t, driver1.FullName, driver2.FullName)
	require.Equal(t, driver1.Email, driver2.Email)
	require.Equal(t, driver1.CabID, driver2.CabID)
	require.WithinDuration(t, driver1.Dob, driver2.Dob, time.Second)
	require.WithinDuration(t, driver1.PasswordChangedAt, driver2.PasswordChangedAt, time.Second)
	require.WithinDuration(t, driver1.CreatedAt, driver2.CreatedAt, time.Second)
}

func TestUpdateDriver(t *testing.T) {
	driver1 := createRandomDriver(t)
	randomCab := createRandomCab(t)

	arg := UpdateDriverParams{
		ID:    driver1.ID,
		CabID: randomCab.ID,
	}

	driver2, err := testQueries.UpdateDriver(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, driver2)

	require.Equal(t, driver1.ID, driver2.ID)
	require.Equal(t, driver1.Username, driver2.Username)
	require.Equal(t, driver1.HashedPassword, driver2.HashedPassword)
	require.Equal(t, driver1.FullName, driver2.FullName)
	require.Equal(t, driver1.Email, driver2.Email)
	require.Equal(t, arg.CabID, driver2.CabID)
	require.WithinDuration(t, driver1.Dob, driver2.Dob, time.Second)
	require.WithinDuration(t, driver1.PasswordChangedAt, driver2.PasswordChangedAt, time.Second)
	require.WithinDuration(t, driver1.CreatedAt, driver2.CreatedAt, time.Second)
}
