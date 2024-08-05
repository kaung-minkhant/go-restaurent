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

type CreateMenuItemParams struct {
	Name        string        `json:"name"`
	Price       float32       `json:"price"`
	Description string        `json:"description"`
	Ingredients string        `json:"ingredients"`
	ImgUrl      string        `json:"img_url"`
	Category    uuid.NullUUID `json:"category"`
	SubCategory uuid.NullUUID `json:"sub_category"`
}

func getMenuItemIDFromPath(r *http.Request) (uuid.UUID, error) {
	idString := chi.URLParam(r, "menuItemId")
	if idString == "" {
		return uuid.UUID{}, fmt.Errorf("require menu item id")
	}
	return uuid.Parse(idString)
}

func handleGetAllMenuItems(w http.ResponseWriter, r *http.Request) error {
	menuItems, err := database.Db.GetAllMenuItems(r.Context())
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(menuItems)
	return nil
}

func handleCreateMenuItem(w http.ResponseWriter, r *http.Request) error {
	newItemParams, err := getRequestBody[CreateMenuItemParams](r)
	if err != nil {
		return err
	}
	menuItem, err := database.Db.CreateMenuItem(r.Context(), models.CreateMenuItemParams{
		ID:          uuid.New(),
		Name:        newItemParams.Name,
		Price:       newItemParams.Price,
		Description: newItemParams.Description,
		Ingredients: newItemParams.Ingredients,
		ImgUrl:      newItemParams.ImgUrl,
		Category:    newItemParams.Category,
		SubCategory: newItemParams.SubCategory,
	})
	if err != nil {
		return err
	}
	return writeJson(w, http.StatusCreated, menuItem)
}

func handleDeleteMenuItem(w http.ResponseWriter, r *http.Request) error {
	menuId, err := getMenuItemIDFromPath(r)
	if err != nil {
		return err
	}
	deletedItem, err := database.Db.DeleteMenuItem(r.Context(), menuId)
	if err != nil {
		return err
	}
	return writeJson(w, http.StatusAccepted, deletedItem)
}

func handleHardDeleteMenuItem(w http.ResponseWriter, r *http.Request) error {
	menuId, err := getMenuItemIDFromPath(r)
	if err != nil {
		return err
	}
	deletedItem, err := database.Db.HardDeleteMenuItem(r.Context(), menuId)
	if err != nil {
		return err
	}
	return writeJson(w, http.StatusAccepted, deletedItem)
}

func handleUpdateMenuItem(w http.ResponseWriter, r *http.Request) error {
	menuItemId, err := getMenuItemIDFromPath(r)
	if err != nil {
		return err
	}
	body, err := getRequestBody[models.UpdateMenuItemParams](r)
	if err != nil {
		return err
	}
	updatedItem, err := database.Db.UpdateMenuItem(r.Context(), menuItemId, models.UpdateMenuItemParams{
		Name:        body.Name,
		Price:       body.Price,
		Description: body.Description,
		ImgUrl:      body.ImgUrl,
		Category:    body.Category,
		SubCategory: body.SubCategory,
	})
	if err != nil {
		return err
	}
	return writeJson(w, http.StatusAccepted, updatedItem)
}

func handleGetMenuItemById(w http.ResponseWriter, r *http.Request) error {
	id, err := getMenuItemIDFromPath(r)
	if err != nil {
		return err
	}
	menuItem, err := database.Db.GetMenuItemById(r.Context(), id)
	if err != nil {
		return err
	}
	return writeJson(w, http.StatusOK, menuItem)
}
