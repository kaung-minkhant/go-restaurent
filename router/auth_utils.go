package router

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/google/uuid"
	"github.com/kaung-minkhant/go-restaurent/auth"
	"github.com/kaung-minkhant/go-restaurent/database"
	"github.com/kaung-minkhant/go-restaurent/database/models"
	"github.com/kaung-minkhant/go-restaurent/utils"
)

func invalidateTokenFamily(r *http.Request, arg models.InvalidateTokenFamilyParams) error {
	return database.Db.InvalidateTokenFamily(r.Context(), arg)
}

func refreshAccessToken(r *http.Request, oldRefreshToken string, oldAccessToken string) (newAccToken string, newRefToken string, err error) {
	tokens, err := database.Db.GetAuthByTokens(r.Context(), models.GetAuthByTokensParams{
		RefreshToken: oldRefreshToken,
		AccessToken:  oldAccessToken,
	})
	if err != nil {
		fmt.Println("Token invalid: cannot find token")
		return "", "", utils.ReturnAccessDenied()
	}
	if !tokens.Valid {
		fmt.Println("Token invalid")
		if err := invalidateTokenFamily(r, models.InvalidateTokenFamilyParams{
			AccessToken:  oldAccessToken,
			RefreshToken: oldRefreshToken,
		}); err != nil {
			fmt.Println("Cannot invalidate tokens", err)
		}
		return "", "", utils.ReturnAccessDenied()
	}

	newRefToken = generateRefreshToken()
	user, err := getUserFromContext(r)
	if err != nil {
		fmt.Println("User not found")
		return "", "", err
	}
	newAccToken, err = auth.GenerateJWT(user)
	if err != nil {
		fmt.Println("Cannot generate refresh token")
		return "", "", utils.ReturnSomethingWentWrong()
	}

	_, err = database.Db.InvalidateToken(r.Context(), models.InvalidateTokenParams{
		RefreshToken: oldRefreshToken,
		AccessToken:  oldAccessToken,
	})

	if err != nil {
		fmt.Println("Cannot invalidate tokens", err)
		return "", "", utils.ReturnSomethingWentWrong()
	}

	_, err = database.Db.CreateTokens(r.Context(), models.CreateTokensParams{
		ID:           uuid.New(),
		AccessToken:  newAccToken,
		RefreshToken: newRefToken,
		Session:      tokens.Session,
		Valid:        true,
	})
	if err != nil {
		fmt.Println("Cannot set tokens", err)
		return "", "", utils.ReturnSomethingWentWrong()
	}
	return newAccToken, newRefToken, nil
}

func generateRefreshToken() string {
	refreshTokenLength := os.Getenv("REFRESH_TOKEN_LENGTH")
	refreshLength, _ := strconv.Atoi(refreshTokenLength)
	return utils.RandString(refreshLength)
}
