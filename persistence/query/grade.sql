-- name: CreateGrade :one
INSERT INTO "Grade" (
    grade_number, course_name
) VALUES (
    $1, $2
)
RETURNING *;

-- name: DeleteGrade :exec
DELETE FROM "Grade"
WHERE id = $1;

-- name: GetGrade :one
SELECT * FROM "Grade"
WHERE id = $1 LIMIT 1;

-- name: ListGrades :many
SELECT * FROM "Grade"
ORDER BY id;

-- name: ListGradesOfCourse :many
SELECT * FROM "Grade" AS g
WHERE g.course_name = $1
ORDER BY g.id;

-- name: UpdateGradeCourseName :one
UPDATE "Grade"
set course_name = $2
WHERE id = $1
    RETURNING *;

-- name: UpdateGradeNumber :one
UPDATE "Grade"
set grade_number = $2
WHERE id = $1
    RETURNING *;