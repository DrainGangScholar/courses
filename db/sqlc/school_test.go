package db

import (
	"context"
	"main/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomSchool(t *testing.T) School {
	arg := CreateSchoolParams{
		Name:    "Mika",
		Address: "Durlan",
		Type:    "Osnovna",
	}
	s, err := testQueries.CreateSchool(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, s)
	require.NotZero(t, s.ID)
	require.NotZero(t, s.CreatedDate)
	require.Equal(t, s.Name, arg.Name)
	require.Equal(t, s.Address, arg.Address)
	require.Equal(t, s.Type, arg.Type)
	return s
}

func CreateRandomSchool(t *testing.T) School {
	return createRandomSchool(t)
}

func TestCreateSchool(t *testing.T) {
	createRandomSchool(t)
}

func TestGetSchool(t *testing.T) {
	s := createRandomSchool(t)
	school, err := testQueries.GetSchool(context.Background(), s.ID)
	require.NoError(t, err)
	require.NotEmpty(t, school)
	require.Equal(t, school, s)
}

func TestDeleteSchool(t *testing.T) {
	s := createRandomSchool(t)
	err := testQueries.DeleteSchool(context.Background(), s.ID)
	require.NoError(t, err)
	s, err = testQueries.GetSchool(context.Background(), s.ID)
	require.Error(t, err)
	require.Empty(t, s)
}

func TestUpdateSchool(t *testing.T) {
	s := createRandomSchool(t)
	arg := UpdateSchoolParams{
		ID:      s.ID,
		Name:    util.RandomString(6),
		Address: util.RandomString(7),
		Type:    util.RandomString(8),
	}
	uS, err := testQueries.UpdateSchool(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, uS)
	require.Equal(t, s.ID, uS.ID)
	require.NotEqual(t, s.Name, uS.Name)
	require.NotEqual(t, s.Address, uS.Address)
	require.NotEqual(t, s.Type, uS.Type)
}
