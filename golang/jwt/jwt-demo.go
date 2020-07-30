package main

import (
	"fmt"
	"log"
	//"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// Secret key to uniquely sign the token
var key = []byte("secret007")

// TokenClaim / TokenPayload
type TokenClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func main() {
	username := "user1234"
	tokenString := GenerateToken(username)
	fmt.Println(tokenString)

	token, _ := ValidateToken(tokenString)
	// Type cast the Claims to *TokenClaims type
	claim := token.Claims.(*TokenClaims)
	fmt.Println(claim.Username)
}

func GenerateToken(username string) string {
	var tokenClaims = TokenClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			// Enter expiration in milisecond
			ExpiresAt: time.Now().Add(10 * time.Minute).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)
	tokenString, err := token.SignedString(key)
	if err != nil {
		log.Fatal(err)
	}
	return tokenString
}

// ValidateToken validates the token with the secret key and return the object
func ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, 
		func(token *jwt.Token)(interface{}, error) {
			return key, nil
		},
	)

	if err != nil {
		// check if Error is Signature Invalid Error
		if err == jwt.ErrSignatureInvalid {
			//return http.StatusUnauthorized
			return nil, err
		}
		//return http.StatusBadRequest
		return nil, err
	}

	// Validate the token if it expired or not
	if !token.Valid {
		// return the Unauthoried Status for expired token
		//return http.StatusUnauthorized
		return nil, err
	}

	return token, nil
}
