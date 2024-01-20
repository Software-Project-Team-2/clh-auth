package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET_TOKEN"))

func GenerateJWT(id int64, username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       id,
		"username": username,
	})

	tokenString, err := token.SignedString(jwtKey)
	return tokenString, err
}

func ParseUserFromToken(tokenString string) (*jwt.MapClaims, bool) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return jwtKey, nil
	})

	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &claims, true
	} else {
		return nil, false
	}
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
