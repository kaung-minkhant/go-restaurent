package models

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type UpdateCagegoryParams struct {
	Name   *string `json:"name"`
	ImgUrl *string `json:"img_url"`
}

func (q *Queries) UpdateCategory(ctx context.Context, id uuid.UUID, params UpdateCagegoryParams) (MenuCategory, error) {
	updateQuery := "UPDATE menu_categories SET "
	queries := make([]string, 0)
	if params.Name != nil {
		queries = append(queries, fmt.Sprintf("name='%v'", *params.Name))
	}
	if params.ImgUrl != nil {
		queries = append(queries, fmt.Sprintf("img_url='%v'", *params.ImgUrl))
	}
	updateQuery += strings.Join(queries, ",")
	updateQuery += fmt.Sprintf(" WHERE id='%v'", id)
	fmt.Println("Update Category Query", updateQuery)
	row := q.db.QueryRowContext(ctx, updateQuery)
	var i MenuCategory
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.ImgUrl,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}
