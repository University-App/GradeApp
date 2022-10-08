package grade_database

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateStudent(t *testing.T) {
	arg := CreateStudentParams{
		LastName:  "LN1",
		FirstName: "FN1",
	}

	student, err := testQueries.CreateStudent(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, student)

	require.Equal(t, arg.LastName, student.LastName)
	require.Equal(t, arg.FirstName, student.FirstName)

	require.NotZero(t, student.ID)

}
