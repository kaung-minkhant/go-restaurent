package router

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/kaung-minkhant/go-restaurent/auth"
	"github.com/kaung-minkhant/go-restaurent/database"
	"github.com/kaung-minkhant/go-restaurent/database/models"
	"github.com/kaung-minkhant/go-restaurent/utils"
)

type CreateUserParams struct {
	Name       string `json:"name"`
	Password   string `json:"password"`
	RoleId     string `json:"role_id"`
	EmployeeID string `json:"employee_id"`
}

type SignInParams struct {
	EmployeeID string `json:"employee_id"`
	Password   string `json:"password"`
}

type SignInResponse struct {
	AccessToken string `json:"access_token"`
}

func handleSignUp(w http.ResponseWriter, r *http.Request) error {
	body, err := getRequestBody[CreateUserParams](r)
	if err != nil {
		return err
	}
	encryptedPassword, err := utils.EncryptPassword(body.Password)
	if err != nil {
		return err
	}
	user, err := database.Db.CreateUser(r.Context(), models.CreateUserParams{
		ID:                uuid.New(),
		Name:              body.Name,
		EncryptedPassword: encryptedPassword,
		EmployeeID:        body.EmployeeID,
		Role:              body.RoleId,
	})
	if err != nil {
		return err
	}
	return writeJson(w, http.StatusCreated, models.DBUserToUser(user))
}

func handleSignIn(w http.ResponseWriter, r *http.Request) error {
	body, err := getRequestBody[SignInParams](r)
	if err != nil {
		return err
	}
	user, err := database.Db.GetUserByEmployeeId(r.Context(), body.EmployeeID)
	if err != nil {
		return err
	}
	if utils.ComparePassword(user.EncryptedPassword, body.Password) != nil {
		return fmt.Errorf("invalid credentials")
	}
	jwt, err := auth.GenerateJWT(&user)
	if err != nil {
		return err
	}
	refreshToken := generateRefreshToken()

	_, err = database.Db.CreateTokens(r.Context(), models.CreateTokensParams{
		ID:           uuid.New(),
		AccessToken:  jwt,
		RefreshToken: refreshToken,
		Session:      uuid.New(),
		Valid:        true,
	})
	if err != nil {
		return utils.ReturnAccessDenied()
	}
	setRefreshToken(w, refreshToken)
	return writeJson(w, http.StatusOK, SignInResponse{AccessToken: jwt})
}

func handleRefreshToken(w http.ResponseWriter, r *http.Request) error {
	refreshToken, err := r.Cookie("refresh-token")
	if err != nil {
		fmt.Println("Refresh token not found")
		return utils.ReturnAccessDenied()
	}
	accToken, err := getAccessTokenFromContext(r)
	if err != nil {
		fmt.Println("Access token not found")
		return err
	}
	jwt, newRefreshToken, err := refreshAccessToken(r, refreshToken.Value, accToken)
	if err != nil {
		return err
	}
	setRefreshToken(w, newRefreshToken)
	return writeJson(w, http.StatusOK, SignInResponse{AccessToken: jwt})
}
