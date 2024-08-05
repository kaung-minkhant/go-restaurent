package router

import (
	"net/http"

	"github.com/kaung-minkhant/go-restaurent/auth"
	"github.com/kaung-minkhant/go-restaurent/database/models"
)

func getUserFromContext(r *http.Request) (*models.User, error) {
	user, ok := r.Context().Value("ctx-user").(*models.User)
	if !ok {
		return nil, ReturnAccessDenied()
	}
	return user, nil
}

func getAccessTokenFromContext(r *http.Request) (string, error) {
	accToken := r.Context().Value("ctx-access-token").(string)
	if accToken == "" {
		return "", ReturnAccessDenied()
	}
	return accToken, nil

}

func getClaimsFromContext(r *http.Request) (*auth.CustomClaims, error) {
	claims := r.Context().Value("ctx-claims").(*auth.CustomClaims)
	if claims == nil {
		return nil, ReturnAccessDenied()
	}
	return claims, nil

}
