// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: province.sql

package db

import (
	"context"
)

const createProvince = `-- name: CreateProvince :one
INSERT INTO province (
  code,
  name
) VALUES (
  $1, $2
) RETURNING id, code, name, created_at
`

type CreateProvinceParams struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

func (q *Queries) CreateProvince(ctx context.Context, arg CreateProvinceParams) (Province, error) {
	row := q.db.QueryRow(ctx, createProvince, arg.Code, arg.Name)
	var i Province
	err := row.Scan(
		&i.ID,
		&i.Code,
		&i.Name,
		&i.CreatedAt,
	)
	return i, err
}

const getProvince = `-- name: GetProvince :one
SELECT id, code, name, created_at FROM province
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetProvince(ctx context.Context, id int64) (Province, error) {
	row := q.db.QueryRow(ctx, getProvince, id)
	var i Province
	err := row.Scan(
		&i.ID,
		&i.Code,
		&i.Name,
		&i.CreatedAt,
	)
	return i, err
}

const updateProvince = `-- name: UpdateProvince :one
UPDATE province
SET
  code = $2,
  name = $3
WHERE
  id = $1
RETURNING id, code, name, created_at
`

type UpdateProvinceParams struct {
	ID   int64  `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

func (q *Queries) UpdateProvince(ctx context.Context, arg UpdateProvinceParams) (Province, error) {
	row := q.db.QueryRow(ctx, updateProvince, arg.ID, arg.Code, arg.Name)
	var i Province
	err := row.Scan(
		&i.ID,
		&i.Code,
		&i.Name,
		&i.CreatedAt,
	)
	return i, err
}
