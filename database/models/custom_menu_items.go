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

func (q *Queries) UpdateMenuItem(ctx context.Context, id uuid.UUID, params UpdateMenuItemParams) (MenuItem, error) {
	updateQuery := "UPDATE menu_items SET "
	queries := make([]string, 0)
	if params.Name != nil {
		queries = append(queries, fmt.Sprintf("name='%v'", *params.Name))
	}
	if params.Price != nil {
		queries = append(queries, fmt.Sprintf("price='%v'", strconv.FormatFloat(float64(*params.Price), 'g', -1, 32)))
	}
	if params.Description != nil {
		queries = append(queries, fmt.Sprintf("description='%v'", *params.Description))
	}
	if params.Ingredients != nil {
		queries = append(queries, fmt.Sprintf("ingredients='%v'", *params.Ingredients))
	}
	if params.ImgUrl != nil {
		queries = append(queries, fmt.Sprintf("img_url='%v'", *params.ImgUrl))
	}
	if params.Category != nil {
		queries = append(queries, fmt.Sprintf("category='%v'", *params.Category))
	}
	if params.SubCategory != nil {
		queries = append(queries, fmt.Sprintf("sub_category='%v'", *params.SubCategory))
	}
	updateQuery += strings.Join(queries, ",")
	updateQuery += fmt.Sprintf(" WHERE id='%v'", id)
	updateQuery += " RETURNING *;"
	fmt.Println("Update menu item query", updateQuery)
	row := q.db.QueryRowContext(ctx, updateQuery)
	var i MenuItem
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Price,
		&i.Description,
		&i.Ingredients,
		&i.ImgUrl,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.Category,
		&i.SubCategory,
	)
	return i, err
}
