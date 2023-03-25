package db

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"uber/util"
)

func createRandomCab(t *testing.T) Cab {
	randomID, err := uuid.NewRandom()
	require.NoError(t, err)
	randomRegNo, err := uuid.NewRandom()
	require.NoError(t, err)

	arg := CreateCabsParams{
		ID:      randomID,
		CabType: util.X,
		RegNo:   randomRegNo,
	}

	cab, err := testQueries.CreateCabs(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, arg.ID, cab.ID)
	require.Equal(t, arg.CabType, cab.CabType)
	require.Equal(t, arg.RegNo, cab.RegNo)
	require.NotZero(t, cab.CreatedAt)
	return cab
}

func TestCreateCab(t *testing.T) {
	createRandomCab(t)
}

func TestGetCab(t *testing.T) {
	cab1 := createRandomCab(t)
	cab2, err := testQueries.GetCab(context.Background(), cab1.ID)
	require.NoError(t, err)

	require.Equal(t, cab1.ID, cab2.ID)
	require.Equal(t, cab1.CabType, cab2.CabType)
	require.Equal(t, cab1.RegNo, cab2.RegNo)
	require.NotZero(t, cab1.CreatedAt, cab2.CreatedAt)
}
