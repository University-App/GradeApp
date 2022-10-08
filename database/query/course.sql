-- name: CreateCourse :one
INSERT INTO "Course" (
    course_name
) VALUES (
    $1
)
RETURNING *;

-- name: DeleteCourse :exec
DELETE FROM "Course"
WHERE id = $1;

-- name: GetCourse :one
SELECT * FROM "Course"
WHERE id = $1 LIMIT 1;

-- name: ListCourses :many
SELECT * FROM "Course"
ORDER BY id;

-- name: UpdateCourseName :one
UPDATE "Course"
set course_name = $2
WHERE id = $1
    RETURNING *;