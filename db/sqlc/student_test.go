package db

import (
	"context"
	"main/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomStudent(t *testing.T) Student {
	school := createRandomSchool(t)
	arg := CreateStudentParams{
		Name:     util.RandomString(7),
		Password: util.RandomString(12),
		SchoolID: school.ID,
	}
	student, err := testQueries.CreateStudent(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, student)
	require.Equal(t, arg.Name, student.Name)
	require.Equal(t, arg.Password, student.Password)
	require.Equal(t, arg.SchoolID, student.SchoolID)
	require.NotZero(t, student.CreatedDate)
	return student
}

func CreateRandomStudent(t *testing.T) Student {
	return createRandomStudent(t)
}

func TestCreateStudent(t *testing.T) {
	createRandomStudent(t)
}

func TestGetStudent(t *testing.T) {
	s := createRandomStudent(t)
	student, err := testQueries.GetStudent(context.Background(), s.ID)
	require.NoError(t, err)
	require.NotEmpty(t, student)
	require.Equal(t, s.Name, s.Name)
	require.Equal(t, s.Password, s.Password)
	require.Equal(t, s.SchoolID, s.SchoolID)
}

func TestDeleteStudent(t *testing.T) {
	s := createRandomStudent(t)
	err := testQueries.DeleteStudent(context.Background(), s.ID)
	require.NoError(t, err)
	_, err = testQueries.GetStudent(context.Background(), s.ID)
	require.Error(t, err)
}

func TestUpdateStudent(t *testing.T) {
	s := createRandomStudent(t)
	arg := UpdateStudentParams{
		ID:       s.ID,
		Name:     util.RandomString(7),
		Password: util.RandomString(12),
	}
	uS, err := testQueries.UpdateStudent(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, uS)
	require.Equal(t, arg.Name, uS.Name)
	require.Equal(t, arg.Password, uS.Password)
}
