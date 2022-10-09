package grade_database

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQueries_CreateStudent(t *testing.T) {
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

	err1 := testQueries.DeleteStudentById(context.Background(), student.ID)
	require.NoError(t, err1)
}

func TestQueries_GetStudent(t *testing.T) {

	studentExpected := Student{
		LastName:  "Djek",
		FirstName: "Pm",
	}

	studentResult, err := testQueries.GetStudent(context.Background(), 1)

	require.NoError(t, err)
	require.NotEmpty(t, studentResult)

	require.Equal(t, studentExpected.LastName, studentResult.LastName)
	require.Equal(t, studentExpected.FirstName, studentResult.FirstName)
}

func TestQueries_GetStudentByFirstLastName(t *testing.T) {

	studentExpected := Student{
		LastName:  "Djek",
		FirstName: "test",
	}

	arg := GetStudentByFirstLastNameParams{
		LastName:  "Djek",
		FirstName: "test",
	}
	studentResult, err := testQueries.GetStudentByFirstLastName(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, studentResult)

	require.Equal(t, studentExpected.LastName, studentResult.LastName)
	require.Equal(t, studentExpected.FirstName, studentResult.FirstName)
}

func TestQueries_ListStudents(t *testing.T) {

	studentsResult, err := testQueries.ListStudents(context.Background())

	require.NoError(t, err)
	require.NotEmpty(t, studentsResult)

	require.Equal(t, 3, len(studentsResult))
}

func TestQueries_DeleteStudentById(t *testing.T) {
	arg := CreateStudentParams{
		LastName:  "LN1",
		FirstName: "FN1",
	}

	studentCreated, errCreation := testQueries.CreateStudent(context.Background(), arg)
	require.NoError(t, errCreation)
	require.NotEmpty(t, studentCreated)

	errDelete := testQueries.DeleteStudentById(context.Background(), studentCreated.ID)
	require.NoError(t, errDelete)

	studentResult, err1 := testQueries.GetStudent(context.Background(), studentCreated.ID)
	require.Error(t, err1)
	require.EqualError(t, err1, sql.ErrNoRows.Error())
	require.Empty(t, studentResult)

}

func TestQueries_DeleteStudentByFirstLastName(t *testing.T) {
	arg := GetStudentByFirstLastNameParams{
		LastName:  "Djek",
		FirstName: "test1",
	}

	deleteArg := DeleteStudentByFirstLastNameParams{
		LastName:  "Djek",
		FirstName: "test1",
	}

	createArg := CreateStudentParams{
		LastName:  "Djek",
		FirstName: "test1",
	}

	errDelete := testQueries.DeleteStudentByFirstLastName(context.Background(), deleteArg)
	require.NoError(t, errDelete)

	studentGot, errGet := testQueries.GetStudentByFirstLastName(context.Background(), arg)
	require.Error(t, errGet)
	require.EqualError(t, errGet, sql.ErrNoRows.Error())
	require.Empty(t, studentGot)

	studentRecreated, err1 := testQueries.CreateStudent(context.Background(), createArg)
	require.NoError(t, err1)
	require.NotEmpty(t, studentRecreated)

}

func TestQueries_UpdateStudentById(t *testing.T) {

	argGet := GetStudentByFirstLastNameParams{
		LastName:  "Djek",
		FirstName: "test",
	}

	studentGot, errGet := testQueries.GetStudentByFirstLastName(context.Background(), argGet)
	require.NoError(t, errGet)
	require.NotEmpty(t, studentGot)

	arg := UpdateStudentByIdParams{
		ID:        studentGot.ID,
		FirstName: "Test",
		LastName:  "Djek",
	}

	studentUpdated, errUpdate := testQueries.UpdateStudentById(context.Background(), arg)
	require.NoError(t, errUpdate)
	require.NotEmpty(t, studentUpdated)

	require.Equal(t, arg.LastName, studentUpdated.LastName)
	require.Equal(t, arg.FirstName, studentUpdated.FirstName)

	argRollback := UpdateStudentByIdParams{
		ID:        studentUpdated.ID,
		FirstName: "test",
		LastName:  "Djek",
	}

	studentRollback, errRollback := testQueries.UpdateStudentById(context.Background(), argRollback)
	require.NoError(t, errRollback)
	require.NotEmpty(t, studentRollback)
}

func TestQueries_UpdateStudentByFirstLastName(t *testing.T) {

	argGet := GetStudentByFirstLastNameParams{
		LastName:  "Djek",
		FirstName: "test",
	}

	studentGot, errGet := testQueries.GetStudentByFirstLastName(context.Background(), argGet)
	require.NoError(t, errGet)
	require.NotEmpty(t, studentGot)

	arg := UpdateStudentByFirstLastNameParams{
		FirstName:   "test",
		LastName:    "Djek",
		FirstName_2: "tes1",
		LastName_2:  "D",
	}

	studentUpdated, errUpdate := testQueries.UpdateStudentByFirstLastName(context.Background(), arg)
	require.NoError(t, errUpdate)
	require.NotEmpty(t, studentUpdated)

	require.Equal(t, arg.LastName_2, studentUpdated.LastName)
	require.Equal(t, arg.FirstName_2, studentUpdated.FirstName)

	argRollback := UpdateStudentByFirstLastNameParams{
		FirstName:   "tes1",
		LastName:    "D",
		FirstName_2: "test",
		LastName_2:  "Djek",
	}

	studentRollback, errRollback := testQueries.UpdateStudentByFirstLastName(context.Background(), argRollback)
	require.NoError(t, errRollback)
	require.NotEmpty(t, studentRollback)
}
