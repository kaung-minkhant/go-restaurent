package models

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

type UpdateMenuItemParams struct {
	Name        *string        `json:"name"`
	Price       *float32       `json:"price"`
	Description *string        `json:"description"`
	Ingredients *string        `json:"ingredients"`
	ImgUrl      *string        `json:"image_url"`
	Category    *uuid.NullUUID `json:"category"`
	SubCategory *uuid.NullUUID `json:"sub_category"`
}

func (q *Queries) UpdateMenuItem(ctx context.Context, id uuid.UUID, params UpdateMenuItemParams) {
	updateQuery := "UPDATE menu_items SET "
	queries := make([]string, 0)
	if params.Name != nil {
		queries = append(queries, "name="+*params.Name)
	}
	if params.Price != nil {
		queries = append(queries, "price="+strconv.FormatFloat(float64(*params.Price), 'g', -1, 32))
	}
	if params.Description != nil {
		queries = append(queries, "description="+*params.Description)
	}
	if params.Ingredients != nil {
		queries = append(queries, "ingredients="+*params.Ingredients)
	}
	if params.ImgUrl != nil {
		queries = append(queries, "image_url="+*params.ImgUrl)
	}
	if params.Category != nil {
		queries = append(queries, "category="+(*params.Category).UUID.String())
	}
	if params.SubCategory != nil {
		queries = append(queries, "sub_category="+(*params.SubCategory).UUID.String())
	}
	updateQuery += strings.Join(queries, ",")
	updateQuery += " WHERE id = " + id.String()
	fmt.Println("Update menu item query", updateQuery)
}
