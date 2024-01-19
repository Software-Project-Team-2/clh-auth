package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"os"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET_TOKEN"))

func GenerateJWT(id uint, username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       id,
		"username": username,
	})

	tokenString, err := token.SignedString(jwtKey)
	return tokenString, err
}

func ValidateToken(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return false
	}

	return token.Valid
}
