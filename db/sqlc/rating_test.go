package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func createRandomRating(t *testing.T) Rating {
	randomUUID, err := uuid.NewRandom()
	require.NoError(t, err)

	trip := createRandomTrip(t)
	arg := CreateRatingParams{
		ID:         randomUUID,
		CustomerID: trip.CustomerID,
		DriverID:   trip.DriverID.UUID,
		TripID:     trip.ID,
		Rating:     int64(5),
		Feedback:   sql.NullString{Valid: false},
	}

	rating, err := testQueries.CreateRating(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, arg.ID, rating.ID)
	require.Equal(t, arg.CustomerID, rating.CustomerID)
	require.Equal(t, arg.DriverID, rating.DriverID)
	require.Equal(t, arg.TripID, rating.TripID)
	require.Equal(t, arg.Rating, rating.Rating)
	require.Equal(t, arg.Feedback, rating.Feedback)

	return rating
}

func TestCreateRating(t *testing.T) {
	createRandomRating(t)
}

func TestGetRating(t *testing.T) {
	rating1 := createRandomRating(t)
	arg := GetRatingParams{
		CustomerID: rating1.CustomerID,
		DriverID:   rating1.DriverID,
		TripID:     rating1.TripID,
	}

	rating2, err := testQueries.GetRating(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, rating1.ID, rating2.ID)
	require.Equal(t, rating1.CustomerID, rating2.CustomerID)
	require.Equal(t, rating1.DriverID, rating2.DriverID)
	require.Equal(t, rating1.TripID, rating2.TripID)
	require.Equal(t, rating1.Rating, rating2.Rating)
	require.Equal(t, rating1.Feedback, rating2.Feedback)
}
