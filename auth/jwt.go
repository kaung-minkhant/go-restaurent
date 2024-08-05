package auth

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kaung-minkhant/go-restaurent/database/models"
	"github.com/kaung-minkhant/go-restaurent/utils"
)

type CustomClaims struct {
	EmployeeID   string `json:"employee_id"`
	EmployeeRole string `json:"employee_role"`
	jwt.RegisteredClaims
}

func GenerateJWT(user *models.User) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	expiryString := os.Getenv("ACC_TOKEN_EXPIRY")
	expiry, _ := strconv.Atoi(expiryString)
	claims := &CustomClaims{
		user.EmployeeID,
		user.Role,
		jwt.RegisteredClaims{
			Issuer:    "Go Restaurent",
			Subject:   "jwt",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * time.Duration(expiry))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", fmt.Errorf("cannot generate jwt")
	}
	return ss, nil
}

func ValidateJWT(jwtToken string) (*CustomClaims, error) {
	secret := os.Getenv("JWT_SECRET")
	token, err := jwt.ParseWithClaims(jwtToken, &CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, utils.ReturnAccessDenied()
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return nil, utils.ReturnAccessDenied()
	}
	return claims, nil
}

func GetClaimsWithoutValidation(jwtToken string) (*CustomClaims, error) {
	secret := os.Getenv("JWT_SECRET")
	token, err := jwt.ParseWithClaims(jwtToken, &CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, utils.ReturnAccessDenied()
		}
		return []byte(secret), nil
	}, jwt.WithoutClaimsValidation())
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return nil, utils.ReturnAccessDenied()
	}
	return claims, nil

}
