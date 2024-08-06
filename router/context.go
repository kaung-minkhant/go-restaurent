package router

import (
	"net/http"

	"github.com/kaung-minkhant/go-restaurent/auth"
	"github.com/kaung-minkhant/go-restaurent/database/models"
	"github.com/kaung-minkhant/go-restaurent/utils"
)

type contextKey string

var ctxUserKey = contextKey("ctx-user")
var ctxClaimsKey = contextKey("ctx-claims")
var ctxAccTokenKey = contextKey("ctx-access-token")
var ctxRefTokenKey = contextKey("ctx-refresh-token")

func getUserFromContext(r *http.Request) (*models.User, error) {
	user, ok := r.Context().Value(ctxUserKey).(*models.User)
	if !ok {
		return nil, utils.ReturnAccessDenied()
	}
	return user, nil
}

func getAccessTokenFromContext(r *http.Request) (string, error) {
	accToken := r.Context().Value(ctxAccTokenKey).(string)
	if accToken == "" {
		return "", utils.ReturnAccessDenied()
	}
	return accToken, nil

}

func getClaimsFromContext(r *http.Request) (*auth.CustomClaims, error) {
	claims := r.Context().Value(ctxClaimsKey).(*auth.CustomClaims)
	if claims == nil {
		return nil, utils.ReturnAccessDenied()
	}
	return claims, nil
}

func getRefreshTokenFromContext(r *http.Request) (string, error) {
	refToken := r.Context().Value(ctxRefTokenKey).(string)
	if refToken == "" {
		return "", utils.ReturnAccessDenied()
	}
	return refToken, nil
}
