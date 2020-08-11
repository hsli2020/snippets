package utils

import (
	"e-food/constants"
	"github.com/dgrijalva/jwt-go"
	"time"
)


func GenerateJWT(userEmail, fname, lname string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = userEmail
	claims["fname"] = fname
	claims["lname"] = lname
	claims["exp"] = time.Now().Add(time.Minute * 300).Unix()

	tokenString, err := token.SignedString(constants.MySecretKeyForJWT)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

package utils

import (
	"e-food/constants"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GenerateJWT(userEmail, fname, lname string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = userEmail
	claims["fname"] = fname
	claims["lname"] = lname
	claims["exp"] = time.Now().Add(time.Minute * 300).Unix()
	/*
	 Please note that in real world, we need to move "some_secret_key_val_123123" into something like
	 "secret.json" file of Kubernetes etc
	*/ 
	tokenString, err := token.SignedString("some_secret_key_val_123123")
	if err != nil {
		return "", err
	}
	return tokenString, nil
}