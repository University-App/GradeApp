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

type StudentCourseAverageTxParams struct {
	CourseName       string `json:"courseName"`
	StudentFirstName string `json:"studentFirstName"`
	StudentLastName  string `json:"studentLastName"`
}

type StudentCourseAverageTxResult struct {
	ID          int32
	CourseName  string `json:"courseName"`
	Average     int32  `json:"average"`
	StudentName string `json:"studentName"`
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

func (store *Store) StudentCourseAverageTx(ctx context.Context, arg StudentCourseAverageTxParams) (StudentCourseAverageTxResult, error) {

	var result StudentCourseAverageTxResult

	err := store.execTx(ctx, func(queries *Queries) error {

		student, errGetStudent := queries.GetStudentByFirstLastName(ctx, GetStudentByFirstLastNameParams{
			LastName:  arg.StudentLastName,
			FirstName: arg.StudentFirstName,
		})
		if errGetStudent != nil {
			return errGetStudent
		}

		studentlistGradesCourse, err := queries.StudentGradesCourse(ctx, StudentGradesCourseParams{CourseName: arg.CourseName, StudentID: sql.NullInt32{Int32: student.ID, Valid: true}})
		if err != nil {
			return err
		}
		for _, grade := range studentlistGradesCourse {
			result.Average += grade.GradeNumber
		}
		result.Average /= int32(len(studentlistGradesCourse))
		result.CourseName = arg.CourseName
		result.StudentName = arg.StudentFirstName + " - " + arg.StudentLastName

		studentCourseAverage, err := queries.CreateStudentCourseAverage(ctx, CreateStudentCourseAverageParams{
			CourseName:  result.CourseName,
			Average:     result.Average,
			StudentName: result.StudentName,
		})
		if err != nil {
			return err
		}

		result.ID = studentCourseAverage.ID

		return nil
	})

	return result, err
}
