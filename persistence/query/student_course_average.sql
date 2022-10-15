-- name: CreateStudentCourseAverage :one
INSERT INTO "StudentCourseAverage" (
    course_name, student_name, average
) VALUES (
    $1, $2, $3
)
RETURNING *;

-- name: ListStudentCourseAverages :many
SELECT * FROM "StudentCourseAverage"
ORDER BY id;

-- name: DeleteStudentCourseAverage :exec
DELETE FROM "StudentCourseAverage"
WHERE id = $1;