// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: roles.sql

package models

import (
	"context"
)

const createRole = `-- name: CreateRole :one
INSERT INTO roles (id, created_at, updated_at, deleted_at)
VALUES ($1,  NOW() AT TIME ZONE 'utc', NOW() AT TIME ZONE 'utc', NULL)
RETURNING id, created_at, updated_at, deleted_at
`

func (q *Queries) CreateRole(ctx context.Context, id string) (Role, error) {
	row := q.db.QueryRowContext(ctx, createRole, id)
	var i Role
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const deleteRole = `-- name: DeleteRole :one
UPDATE roles
SET deleted_at = NOW() AT TIME ZONE 'utc'
WHERE id = $1
RETURNING id, created_at, updated_at, deleted_at
`

func (q *Queries) DeleteRole(ctx context.Context, id string) (Role, error) {
	row := q.db.QueryRowContext(ctx, deleteRole, id)
	var i Role
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getAllRoles = `-- name: GetAllRoles :many
SELECT id, created_at, updated_at, deleted_at FROM roles
`

func (q *Queries) GetAllRoles(ctx context.Context) ([]Role, error) {
	rows, err := q.db.QueryContext(ctx, getAllRoles)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Role
	for rows.Next() {
		var i Role
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getRoleById = `-- name: GetRoleById :one
SELECT id, created_at, updated_at, deleted_at FROM roles
WHERE id = $1
`

func (q *Queries) GetRoleById(ctx context.Context, id string) (Role, error) {
	row := q.db.QueryRowContext(ctx, getRoleById, id)
	var i Role
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}
