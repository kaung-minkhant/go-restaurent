package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/kaung-minkhant/go-restaurent/database"
	"github.com/kaung-minkhant/go-restaurent/database/models"
)

type CreateSubCategoryParams struct {
	Name     string    `json:"name"`
	ImgUrl   string    `json:"img_url"`
	Category uuid.UUID `json:"category"`
}

func getSubCategoryIdFromParam(r *http.Request) (uuid.UUID, error) {
	idString := chi.URLParam(r, "subCategoryId")
	if idString == "" {
		return uuid.UUID{}, fmt.Errorf("sub category id is required")
	}
	return uuid.Parse(idString)
}

func handleGetAllSubCategory(w http.ResponseWriter, r *http.Request) error {
	subCategory, err := database.Db.GetAllSubCategory(r.Context())
	if err != nil {
		return err
	}

	return writeJson(w, http.StatusOK, subCategory)
}

func handleCreateSubCategory(w http.ResponseWriter, r *http.Request) error {
	body, err := getRequestBody[CreateSubCategoryParams](r)
	if err != nil {
		return err
	}
	subCat, err := database.Db.CreateSubCategory(r.Context(), models.CreateSubCategoryParams{
		ID:       uuid.New(),
		Name:     body.Name,
		ImgUrl:   body.ImgUrl,
		Category: body.Category,
	})
	if err != nil {
		return nil
	}
	return writeJson(w, http.StatusCreated, subCat)
}

func handleDeleteSubCategory(w http.ResponseWriter, r *http.Request) error {
	id, err := getSubCategoryIdFromParam(r)
	if err != nil {
		return err
	}
	deleteSubCat, err := database.Db.DeleteSubCategory(r.Context(), id)
	if err != nil {
		return err
	}
	return writeJson(w, http.StatusAccepted, deleteSubCat)
}

func handleUpdateSubCategory(w http.ResponseWriter, r *http.Request) error {
	id, err := getSubCategoryIdFromParam(r)
	if err != nil {
		return err
	}
	body, err := getRequestBody[models.UpdateSubCagegoryParams](r)
	if err != nil {
		return err
	}
	updated, err := database.Db.UpdateSubCategory(r.Context(), id, *body)
	if err != nil {
		return err
	}
	return writeJson(w, http.StatusAccepted, updated)
}
