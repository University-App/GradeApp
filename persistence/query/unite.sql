-- name: CreateUnite :one
INSERT INTO "Unite" (
    unite_name
) VALUES (
    $1
)
RETURNING *;

-- name: DeleteUnite :exec
DELETE FROM "Unite"
WHERE id = $1;

-- name: GetUnite :one
SELECT * FROM "Unite"
WHERE id = $1 LIMIT 1;

-- name: ListUnites :many
SELECT * FROM "Unite"
ORDER BY id;

-- name: UpdateUniteName :one
UPDATE "Unite"
set unite_name = $2
WHERE id = $1
    RETURNING *;