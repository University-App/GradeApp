package grade_database

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (store *Store) execTx(ctx context.Context, fn func(queries *Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %w, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

type CourseAverageTxParams struct {
	CourseName string `json:"courseName"`
}

type CourseAverageTxResult struct {
	ID         int32
	CourseName string `json:"courseName"`
	Average    int32  `json:"average"`
}

func (store *Store) CourseAverageTx(ctx context.Context, arg CourseAverageTxParams) (CourseAverageTxResult, error) {

	var result CourseAverageTxResult

	err := store.execTx(ctx, func(queries *Queries) error {

		listGrades, err := queries.ListGradesOfCourse(ctx, arg.CourseName)
		if err != nil {
			return err
		}
		for _, grade := range listGrades {
			result.Average += grade.GradeNumber
		}
		result.Average /= int32(len(listGrades))
		result.CourseName = arg.CourseName

		courseAverage, err := queries.CreateCourseAverage(ctx, CreateCourseAverageParams{
			CourseName: result.CourseName,
			Average:    result.Average,
		})
		if err != nil {
			return err
		}

		result.ID = courseAverage.ID

		return nil
	})

	return result, err
}
