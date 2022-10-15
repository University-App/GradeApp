-- name: CreateStudent :one
INSERT INTO "Student" (
    last_name, first_name
) VALUES (
    $1, $2
)
RETURNING *;

-- name: DeleteStudentById :exec
DELETE FROM "Student"
WHERE id = $1;

-- name: DeleteStudentByFirstLastName :exec
DELETE FROM "Student"
WHERE first_name = $1 AND last_name = $2;

-- name: GetStudent :one
SELECT * FROM "Student"
WHERE id = $1 LIMIT 1;

-- name: GetStudentByFirstLastName :one
SELECT * FROM "Student"
WHERE first_name = $1 AND last_name = $2 LIMIT 1;

-- name: ListStudents :many
SELECT * FROM "Student"
ORDER BY id;

-- name: UpdateStudentById :one
UPDATE "Student"
set first_name = $2,
    last_name = $3
WHERE id = $1
    RETURNING *;

-- name: UpdateStudentByFirstLastName :one
UPDATE "Student"
set first_name = $3,
    last_name = $4
WHERE first_name = $1 AND last_name = $2
    RETURNING *;