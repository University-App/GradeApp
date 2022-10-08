-- name: CreateStudent :one
INSERT INTO "Student" (
    last_name, first_name
) VALUES (
    $1, $2
)
RETURNING *;

-- name: DeleteStudent :exec
DELETE FROM "Student"
WHERE id = $1;

-- name: GetStudent :one
SELECT * FROM "Student"
WHERE id = $1 LIMIT 1;

-- name: ListStudents :many
SELECT * FROM "Student"
ORDER BY id;

-- name: UpdateStudent :one
UPDATE "Student"
set first_name = $2,
    last_name = $3
WHERE id = $1
    RETURNING *;