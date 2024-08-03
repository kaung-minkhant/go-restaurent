package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/kaung-minkhant/go-restaurent/database"
	"github.com/kaung-minkhant/go-restaurent/database/models"
)

type CreateCategoryParams struct {
	Name   string `json:"name"`
	ImgUrl string `json:"img_url"`
}

func getRequestBody[T any](r *http.Request) (*T, error) {
	var body T
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}
	return &body, nil
}

func getCategoryIdFromParam(r *http.Request) (uuid.UUID, error) {
	idString := chi.URLParam(r, "categoryId")
	if idString == "" {
		return uuid.UUID{}, fmt.Errorf("category id is required")
	}
	return uuid.Parse(idString)
}

func handleGetAllCategory(w http.ResponseWriter, r *http.Request) error {
	categories, err := database.Db.GetAllCategory(r.Context())
	if err != nil {
		return err
	}

	return writeJson(w, http.StatusOK, categories)
}

func handleCreateCategory(w http.ResponseWriter, r *http.Request) error {
	body, err := getRequestBody[CreateCategoryParams](r)
	if err != nil {
		return err
	}
	cat, err := database.Db.CreateCategory(r.Context(), models.CreateCategoryParams{
		ID:     uuid.New(),
		Name:   body.Name,
		ImgUrl: body.ImgUrl,
	})
	if err != nil {
		return nil
	}
	return writeJson(w, http.StatusCreated, cat)
}

func handleDeleteCategory(w http.ResponseWriter, r *http.Request) error {
	id, err := getCategoryIdFromParam(r)
	if err != nil {
		return err
	}
	deletedCat, err := database.Db.DeleteCategory(r.Context(), id)
	if err != nil {
		return err
	}
	return writeJson(w, http.StatusAccepted, deletedCat)
}

func handleUpdateCategory(w http.ResponseWriter, r *http.Request) error {
	id, err := getCategoryIdFromParam(r)
	if err != nil {
		return err
	}
	body, err := getRequestBody[models.UpdateCagegoryParams](r)
	if err != nil {
		return err
	}
	updated, err := database.Db.UpdateCategory(r.Context(), id, *body)
	if err != nil {
		return err
	}
	return writeJson(w, http.StatusAccepted, updated)
}
