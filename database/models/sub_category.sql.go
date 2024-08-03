// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: sub_category.sql

package models

import (
	"context"

	"github.com/google/uuid"
)

const createSubCategory = `-- name: CreateSubCategory :one
INSERT INTO menu_sub_categories (id, name, img_url, category, created_at, updated_at, deleted_at)
VALUES ($1, $2, $3, $4, NOW() AT TIME ZONE 'utc', NOW() AT TIME ZONE 'utc', NULL)
RETURNING id, name, img_url, category, created_at, updated_at, deleted_at
`

type CreateSubCategoryParams struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	ImgUrl   string    `json:"img_url"`
	Category uuid.UUID `json:"category"`
}

func (q *Queries) CreateSubCategory(ctx context.Context, arg CreateSubCategoryParams) (MenuSubCategory, error) {
	row := q.db.QueryRowContext(ctx, createSubCategory,
		arg.ID,
		arg.Name,
		arg.ImgUrl,
		arg.Category,
	)
	var i MenuSubCategory
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.ImgUrl,
		&i.Category,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const deleteSubCategory = `-- name: DeleteSubCategory :one
UPDATE menu_sub_categories
SET deleted_at = NOW()
WHERE id = $1
RETURNING id, name, img_url, category, created_at, updated_at, deleted_at
`

func (q *Queries) DeleteSubCategory(ctx context.Context, id uuid.UUID) (MenuSubCategory, error) {
	row := q.db.QueryRowContext(ctx, deleteSubCategory, id)
	var i MenuSubCategory
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.ImgUrl,
		&i.Category,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getAllSubCategory = `-- name: GetAllSubCategory :many
SELECT id, name, img_url, category, created_at, updated_at, deleted_at FROM menu_sub_categories
`

func (q *Queries) GetAllSubCategory(ctx context.Context) ([]MenuSubCategory, error) {
	rows, err := q.db.QueryContext(ctx, getAllSubCategory)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []MenuSubCategory
	for rows.Next() {
		var i MenuSubCategory
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.ImgUrl,
			&i.Category,
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

const getSubCategoryById = `-- name: GetSubCategoryById :one
SELECT id, name, img_url, category, created_at, updated_at, deleted_at FROM menu_sub_categories
WHERE id = $1
`

func (q *Queries) GetSubCategoryById(ctx context.Context, id uuid.UUID) (MenuSubCategory, error) {
	row := q.db.QueryRowContext(ctx, getSubCategoryById, id)
	var i MenuSubCategory
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.ImgUrl,
		&i.Category,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}
