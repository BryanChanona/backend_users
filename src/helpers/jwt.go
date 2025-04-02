package helpers

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
	IdUser int `json:"id_user"`
	jwt.StandardClaims
}

func GenerateJWT(IdUser int) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		IdUser: IdUser,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
