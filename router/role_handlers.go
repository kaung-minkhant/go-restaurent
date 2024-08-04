package router

import (
	"net/http"

	"github.com/kaung-minkhant/go-restaurent/database"
)

type CreateRoleParams struct {
	RoleId string `json:"role_id"`
}

func handleCreateRole(w http.ResponseWriter, r *http.Request) error {
	body, err := getRequestBody[CreateRoleParams](r)
	if err != nil {
		return err
	}
	role, err := database.Db.CreateRole(r.Context(), body.RoleId)
	if err != nil {
		return err
	}

	return writeJson(w, http.StatusCreated, role)
}
