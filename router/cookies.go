package router

import (
	"net/http"

	"github.com/kaung-minkhant/go-restaurent/utils"
)

func setRefreshTokenCookie(w http.ResponseWriter, refreshToken string) {
	http.SetCookie(w, &http.Cookie{
		Name:     "refresh-token",
		Value:    refreshToken,
		HttpOnly: true,
	})
}
func getRefreshTokenFromCookie(r *http.Request) (string, error) {
	refreshToken, err := r.Cookie("refresh-token")
	if err != nil {
		return "", utils.ReturnAccessDenied()
	}
	return refreshToken.Value, nil

}
