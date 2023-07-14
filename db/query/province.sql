-- name: CreateProvince :one
INSERT INTO province (
  code,
  name
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetProvince :one
SELECT * FROM province
WHERE id = $1 LIMIT 1;

-- name: UpdateProvince :one
UPDATE province
SET
  code = $2,
  name = $3
WHERE
  id = $1
RETURNING *;
