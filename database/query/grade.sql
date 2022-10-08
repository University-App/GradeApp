-- name: CreateGrade :one
INSERT INTO "Grade" (
    number, course_name
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

-- name: UpdateGradeCourseName :one
UPDATE "Grade"
set course_name = $2
WHERE id = $1
    RETURNING *;

-- name: UpdateGradeNumber :one
UPDATE "Grade"
set number = $2
WHERE id = $1
    RETURNING *;