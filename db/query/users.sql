-- name: CreateUser :one
INSERT INTO users (
  username,
  prov_id
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: UpdateUser :one
UPDATE users
SET
  username = COALESCE(sqlc.narg(username), username),
  prov_id = COALESCE(sqlc.narg(prov_id), prov_id)
WHERE
  id = sqlc.arg(id)
RETURNING *;
