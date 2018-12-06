package util

import (
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

// NewUUID generate and return a UUID string
func NewUUID() string {
	u := uuid.Must(uuid.NewV4()).String()
	return strings.Replace(u, "-", "", -1)
}

// GenJWT generate and return a JWT token
func GenJWT(claims *jwt.StandardClaims, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(String2Bytes(secret))
}

// EncryptPassword encrypt string by bcrypt
// Use to encrypt password
func EncryptPassword(password string) (string, error) {
	retBytes, err := bcrypt.GenerateFromPassword(String2Bytes(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(retBytes[:]), err
}

// ComparePassword use to compare a string with a hashed password
func ComparePassword(hasedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword(String2Bytes(hasedPassword), String2Bytes(password))
}

// String2Bytes convert a string to []byte
func String2Bytes(str string) []byte {
	return []byte(str)
}
