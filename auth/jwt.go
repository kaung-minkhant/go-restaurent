package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kaung-minkhant/go-restaurent/database/models"
)

type CustomClaims struct {
	EmployeeID   string `json:"employee_id"`
	EmployeeRole string `json:"employee_role"`
	jwt.RegisteredClaims
}

// TODO: properly store JWT secret
var secret string = "secret"

func GenerateJWT(user *models.User) (string, error) {
	claims := &CustomClaims{
		user.EmployeeID,
		user.Role,
		jwt.RegisteredClaims{
			Issuer:    "Go Restaurent",
			Subject:   "jwt",
			Audience:  []string{"admin", "user"},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 1)),
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
	token, err := jwt.ParseWithClaims(jwtToken, &CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("access denied")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return nil, fmt.Errorf("access denied")
	}
	return claims, nil
}

func GetClaimsWithoutValidation(jwtToken string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(jwtToken, &CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("access denied")
		}
		return []byte(secret), nil
	}, jwt.WithoutClaimsValidation())
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return nil, fmt.Errorf("access denied")
	}
	return claims, nil

}
