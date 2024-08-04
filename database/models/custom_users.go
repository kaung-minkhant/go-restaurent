package models

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

type UpdateUserParams struct {
	Name       *string `json:"name"`
	Role       *string `json:"role"`
	EmployeeID *string `json:"employee_id"`
}

type CustomUser struct {
	ID         uuid.UUID    `json:"id"`
	Name       string       `json:"name"`
	Role       string       `json:"role"`
	EmployeeID string       `json:"employee_id"`
	CreatedAt  time.Time    `json:"created_at"`
	UpdatedAt  time.Time    `json:"updated_at"`
	DeletedAt  sql.NullTime `json:"deleted_at"`
}

func DBUserToUser(user User) *CustomUser {
	return &CustomUser{
		ID:         user.ID,
		Name:       user.Name,
		Role:       user.Role,
		EmployeeID: user.EmployeeID,
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
		DeletedAt:  user.DeletedAt,
	}
}

func (q *Queries) UpdateUser(ctx context.Context, id uuid.UUID, params UpdateUserParams) (User, error) {
	updateQuery := "UPDATE users SET "
	queries := make([]string, 0)
	if params.Name != nil {
		queries = append(queries, fmt.Sprintf("name='%v'", *params.Name))
	}
	if params.Role != nil {
		queries = append(queries, fmt.Sprintf("role='%v'", *params.Role))
	}
	if params.EmployeeID != nil {
		queries = append(queries, fmt.Sprintf("employee_id='%v'", *params.EmployeeID))
	}
	updateQuery += strings.Join(queries, ",")
	updateQuery += fmt.Sprintf(" WHERE id='%v'", id)
	fmt.Println("Update User Query", updateQuery)
	row := q.db.QueryRowContext(ctx, updateQuery)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.EncryptedPassword,
		&i.Role,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}
