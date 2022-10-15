-- name: CreateCourseAverage :one
INSERT INTO "CourseAverage" (
    course_name, average
) VALUES (
    $1, $2
)
RETURNING *;

-- name: ListCourseAverages :many
SELECT * FROM "CourseAverage"
ORDER BY id;

-- name: DeleteCourseAverage :exec
DELETE FROM "CourseAverage"
WHERE id = $1;