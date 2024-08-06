package router

import (
	"net/http"

	"github.com/kaung-minkhant/go-restaurent/auth"
	"github.com/kaung-minkhant/go-restaurent/database/models"
	"github.com/kaung-minkhant/go-restaurent/utils"
)

func getUserFromContext(r *http.Request) (*models.User, error) {
	user, ok := r.Context().Value("ctx-user").(*models.User)
	if !ok {
		return nil, utils.ReturnAccessDenied()
	}
	return user, nil
}

func getAccessTokenFromContext(r *http.Request) (string, error) {
	accToken := r.Context().Value("ctx-access-token").(string)
	if accToken == "" {
		return "", utils.ReturnAccessDenied()
	}
	return accToken, nil

}

func getClaimsFromContext(r *http.Request) (*auth.CustomClaims, error) {
	claims := r.Context().Value("ctx-claims").(*auth.CustomClaims)
	if claims == nil {
		return nil, utils.ReturnAccessDenied()
	}
	return claims, nil
}

func getRefreshTokenFromContext(r *http.Request) (string, error) {
	refToken := r.Context().Value("ctx-refresh-token").(string)
	if refToken == "" {
		return "", utils.ReturnAccessDenied()
	}
	return refToken, nil
}
