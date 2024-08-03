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

func (q *Queries) UpdateSubCategory(ctx context.Context, id uuid.UUID, params UpdateSubCagegoryParams) {
	updateQuery := "UPDATE menu_sub_categories SET "
	queries := make([]string, 0)
	if params.Name != nil {
		queries = append(queries, "name="+*params.Name)
	}
	if params.ImgUrl != nil {
		queries = append(queries, "img_url="+*params.ImgUrl)
	}
	if params.Category != nil {
		queries = append(queries, "category="+(*params.Category).String())
	}
	updateQuery += strings.Join(queries, ",")
	updateQuery += " WHERE id = " + id.String()
	fmt.Println("Update Sub Category Query", updateQuery)
}
