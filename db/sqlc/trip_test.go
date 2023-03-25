package db

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"uber/util"
)

func createRandomTrip(t *testing.T) Trip {
	randomUUID, err := uuid.NewRandom()
	require.NoError(t, err)

	customer := createRandomCustomer(t)
	driver := createRandomDriver(t)
	driverID := uuid.NullUUID{UUID: driver.ID, Valid: true}

	arg := CreateTripParams{
		ID:         randomUUID,
		CustomerID: customer.ID,
		DriverID:   driverID,
		Status:     util.Creating,
	}

	trip, err := testQueries.CreateTrip(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, arg.ID, trip.ID)
	require.Equal(t, arg.CustomerID, trip.CustomerID)
	require.Equal(t, arg.DriverID, trip.DriverID)
	require.Equal(t, arg.Status, trip.Status)
	require.NotZero(t, trip.CreatedAt)

	return trip
}

func TestCreateTrip(t *testing.T) {
	createRandomTrip(t)
}

func TestGetTrip(t *testing.T) {
	trip1 := createRandomTrip(t)
	arg := GetTripParams{
		CustomerID: trip1.CustomerID,
		DriverID:   trip1.DriverID,
	}

	trip2, err := testQueries.GetTrip(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, trip1.ID, trip2.ID)
	require.Equal(t, trip1.CustomerID, trip2.CustomerID)
	require.Equal(t, trip1.DriverID, trip2.DriverID)
	require.Equal(t, trip1.Status, trip2.Status)
	require.WithinDuration(t, trip1.CreatedAt, trip2.CreatedAt, time.Second)
}

func TestUpdateTrip(t *testing.T) {
	trip1 := createRandomTrip(t)
	arg := UpdateTripParams{
		CustomerID: trip1.CustomerID,
		DriverID:   trip1.DriverID,
		Status:     trip1.Status,
		Status_2:   util.Done,
	}

	trip2, err := testQueries.UpdateTrip(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, trip1.ID, trip2.ID)
	require.Equal(t, trip1.CustomerID, trip2.CustomerID)
	require.Equal(t, trip1.DriverID, trip2.DriverID)
	require.Equal(t, arg.Status_2, trip2.Status)
	require.WithinDuration(t, trip1.CreatedAt, trip2.CreatedAt, time.Second)
}
