// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: grade.sql

package grade_database

import (
	"context"
)

const createGrade = `-- name: CreateGrade :one
INSERT INTO "Grade" (
    grade_number, course_name
) VALUES (
    $1, $2
)
RETURNING id, grade_number, course_name, student_id
`

type CreateGradeParams struct {
	GradeNumber int32
	CourseName  string
}

func (q *Queries) CreateGrade(ctx context.Context, arg CreateGradeParams) (Grade, error) {
	row := q.db.QueryRowContext(ctx, createGrade, arg.GradeNumber, arg.CourseName)
	var i Grade
	err := row.Scan(
		&i.ID,
		&i.GradeNumber,
		&i.CourseName,
		&i.StudentID,
	)
	return i, err
}

const deleteGrade = `-- name: DeleteGrade :exec
DELETE FROM "Grade"
WHERE id = $1
`

func (q *Queries) DeleteGrade(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteGrade, id)
	return err
}

const getGrade = `-- name: GetGrade :one
SELECT id, grade_number, course_name, student_id FROM "Grade"
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetGrade(ctx context.Context, id int32) (Grade, error) {
	row := q.db.QueryRowContext(ctx, getGrade, id)
	var i Grade
	err := row.Scan(
		&i.ID,
		&i.GradeNumber,
		&i.CourseName,
		&i.StudentID,
	)
	return i, err
}

const listGrades = `-- name: ListGrades :many
SELECT id, grade_number, course_name, student_id FROM "Grade"
ORDER BY id
`

func (q *Queries) ListGrades(ctx context.Context) ([]Grade, error) {
	rows, err := q.db.QueryContext(ctx, listGrades)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Grade
	for rows.Next() {
		var i Grade
		if err := rows.Scan(
			&i.ID,
			&i.GradeNumber,
			&i.CourseName,
			&i.StudentID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listGradesOfCourse = `-- name: ListGradesOfCourse :many
SELECT id, grade_number, course_name, student_id FROM "Grade" AS g
WHERE g.course_name = $1
ORDER BY g.id
`

func (q *Queries) ListGradesOfCourse(ctx context.Context, courseName string) ([]Grade, error) {
	rows, err := q.db.QueryContext(ctx, listGradesOfCourse, courseName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Grade
	for rows.Next() {
		var i Grade
		if err := rows.Scan(
			&i.ID,
			&i.GradeNumber,
			&i.CourseName,
			&i.StudentID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateGradeCourseName = `-- name: UpdateGradeCourseName :one
UPDATE "Grade"
set course_name = $2
WHERE id = $1
    RETURNING id, grade_number, course_name, student_id
`

type UpdateGradeCourseNameParams struct {
	ID         int32
	CourseName string
}

func (q *Queries) UpdateGradeCourseName(ctx context.Context, arg UpdateGradeCourseNameParams) (Grade, error) {
	row := q.db.QueryRowContext(ctx, updateGradeCourseName, arg.ID, arg.CourseName)
	var i Grade
	err := row.Scan(
		&i.ID,
		&i.GradeNumber,
		&i.CourseName,
		&i.StudentID,
	)
	return i, err
}

const updateGradeNumber = `-- name: UpdateGradeNumber :one
UPDATE "Grade"
set grade_number = $2
WHERE id = $1
    RETURNING id, grade_number, course_name, student_id
`

type UpdateGradeNumberParams struct {
	ID          int32
	GradeNumber int32
}

func (q *Queries) UpdateGradeNumber(ctx context.Context, arg UpdateGradeNumberParams) (Grade, error) {
	row := q.db.QueryRowContext(ctx, updateGradeNumber, arg.ID, arg.GradeNumber)
	var i Grade
	err := row.Scan(
		&i.ID,
		&i.GradeNumber,
		&i.CourseName,
		&i.StudentID,
	)
	return i, err
}
