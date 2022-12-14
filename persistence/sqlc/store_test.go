package grade_database

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStore_CourseAverageTx(t *testing.T) {
	store := NewStore(testDB)

	// run n conccurent course average transactions
	n := 5

	errs := make(chan error)
	results := make(chan CourseAverageTxResult)

	for i := 0; i < n; i++ {
		go func() {
			result, err := store.CourseAverageTx(context.Background(), CourseAverageTxParams{
				CourseName: "Course1",
			})

			errs <- err
			results <- result
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results

		require.NotEmpty(t, result)
		require.NotZero(t, result.Average)
		require.Equal(t, "Course1", result.CourseName)
	}

	listeCourseAverages, err := testQueries.ListCourseAverages(context.Background())
	require.NoError(t, err)

	for _, courseAverage := range listeCourseAverages {
		errDelete := testQueries.DeleteCourseAverage(context.Background(), courseAverage.ID)
		require.NoError(t, errDelete)
	}
}

func TestStore_StudentCourseAverageTx(t *testing.T) {
	store := NewStore(testDB)

	// run n conccurent course average transactions
	n := 5

	errs := make(chan error)
	results := make(chan StudentCourseAverageTxResult)

	for i := 0; i < n; i++ {
		go func() {
			result, err := store.StudentCourseAverageTx(context.Background(), StudentCourseAverageTxParams{
				CourseName:       "Course1",
				StudentLastName:  "Djek",
				StudentFirstName: "Pm",
			})
			errs <- err
			results <- result
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results

		require.NotEmpty(t, result)
		require.NotZero(t, result.Average)
		require.Equal(t, "Course1", result.CourseName)
	}

	listeStudentCourseAverages, err := testQueries.ListStudentCourseAverages(context.Background())
	require.NoError(t, err)

	for _, studentCourseAverage := range listeStudentCourseAverages {
		errDelete := testQueries.DeleteStudentCourseAverage(context.Background(), studentCourseAverage.ID)
		require.NoError(t, errDelete)
	}
}
