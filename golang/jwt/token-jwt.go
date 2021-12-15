package helpers

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// CreateToken : create a jwt token
func CreateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"username": username,
		"nbf":      time.Now().Add(time.Second * 3).Unix(),
	})

	return token.SignedString([]byte(os.Getenv("SECRET")))
}

// VerifyToken : verify a token
func VerifyToken(tokenString, username string) bool {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims["username"] == username {
			return true
		}
	}

	return false
}
