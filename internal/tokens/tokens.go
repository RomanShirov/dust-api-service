package tokens

import (
	"github.com/golang-jwt/jwt/v4"
	"os"
)

var jwtSecretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

func GenerateUserToken(username string) string {
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
	})
	token, _ := tokenClaims.SignedString(jwtSecretKey)
	return token
}

func GetUsernameFromToken(token jwt.Token) {}
