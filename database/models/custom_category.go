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

func (q *Queries) UpdateCategory(ctx context.Context, id uuid.UUID, params UpdateCagegoryParams) {
	updateQuery := "UPDATE menu_categories SET "
	queries := make([]string, 0)
	if params.Name != nil {
		queries = append(queries, "name="+*params.Name)
	}
	if params.ImgUrl != nil {
		queries = append(queries, "img_url="+*params.ImgUrl)
	}
	updateQuery += strings.Join(queries, ",")
	updateQuery += " WHERE id = " + id.String()
	fmt.Println("Update Category Query", updateQuery)
}
