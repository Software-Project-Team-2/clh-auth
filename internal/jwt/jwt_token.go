package jwt

import "github.com/dgrijalva/jwt-go"

var jwtKey = []byte("your_secret_key") // Define this key securely

func generateJWT(id uint, username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       id,
		"username": username,
	})

	tokenString, err := token.SignedString(jwtKey)
	return tokenString, err
}

func validateToken(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return false
	}

	return token.Valid
}

