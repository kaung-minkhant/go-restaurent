package models

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type UpdateSubCagegoryParams struct {
	Name     *string    `json:"name"`
	ImgUrl   *string    `json:"img_url"`
	Category *uuid.UUID `json:"category"`
}

func (q *Queries) UpdateSubCategory(ctx context.Context, id uuid.UUID, params UpdateSubCagegoryParams) (MenuSubCategory, error) {
	updateQuery := "UPDATE menu_sub_categories SET "
	queries := make([]string, 0)
	if params.Name != nil {
		queries = append(queries, fmt.Sprintf("name='%v'", *params.Name))
	}
	if params.ImgUrl != nil {
		queries = append(queries, fmt.Sprintf("img_url='%v'", *params.ImgUrl))
	}
	if params.Category != nil {
		queries = append(queries, fmt.Sprintf("category='%v'", *params.Category))
	}
	updateQuery += strings.Join(queries, ",")
	updateQuery += fmt.Sprintf(" WHERE id='%v'", id)
	fmt.Println("Update Sub Category Query", updateQuery)
	row := q.db.QueryRowContext(ctx, createSubCategory)
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
