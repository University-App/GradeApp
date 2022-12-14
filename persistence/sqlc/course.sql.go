// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: course.sql

package grade_database

import (
	"context"
)

const createCourse = `-- name: CreateCourse :one
INSERT INTO "Course" (
    course_name
) VALUES (
    $1
)
RETURNING id, course_name, unite_id
`

func (q *Queries) CreateCourse(ctx context.Context, courseName string) (Course, error) {
	row := q.db.QueryRowContext(ctx, createCourse, courseName)
	var i Course
	err := row.Scan(&i.ID, &i.CourseName, &i.UniteID)
	return i, err
}

const deleteCourse = `-- name: DeleteCourse :exec
DELETE FROM "Course"
WHERE id = $1
`

func (q *Queries) DeleteCourse(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteCourse, id)
	return err
}

const getCourse = `-- name: GetCourse :one
SELECT id, course_name, unite_id FROM "Course"
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetCourse(ctx context.Context, id int32) (Course, error) {
	row := q.db.QueryRowContext(ctx, getCourse, id)
	var i Course
	err := row.Scan(&i.ID, &i.CourseName, &i.UniteID)
	return i, err
}

const listCourses = `-- name: ListCourses :many
SELECT id, course_name, unite_id FROM "Course"
ORDER BY id
`

func (q *Queries) ListCourses(ctx context.Context) ([]Course, error) {
	rows, err := q.db.QueryContext(ctx, listCourses)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Course
	for rows.Next() {
		var i Course
		if err := rows.Scan(&i.ID, &i.CourseName, &i.UniteID); err != nil {
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

const updateCourseName = `-- name: UpdateCourseName :one
UPDATE "Course"
set course_name = $2
WHERE id = $1
    RETURNING id, course_name, unite_id
`

type UpdateCourseNameParams struct {
	ID         int32
	CourseName string
}

func (q *Queries) UpdateCourseName(ctx context.Context, arg UpdateCourseNameParams) (Course, error) {
	row := q.db.QueryRowContext(ctx, updateCourseName, arg.ID, arg.CourseName)
	var i Course
	err := row.Scan(&i.ID, &i.CourseName, &i.UniteID)
	return i, err
}
