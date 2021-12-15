package helpers

import (
	"log"

	"github.com/matthewhartstonge/argon2"
)

// CreateHash : create and return a hash from a password
func CreateHash(password string) string {
	hasher := argon2.DefaultConfig()

	raw, err := hasher.Hash([]byte(password), nil)

	if err != nil {
		log.Fatalf("Error during hashing: %s\n", err.Error())
	}

	return string(raw.Encode())
}

// VerifyHash : verifies hash against password and returns result
func VerifyHash(hash, password string) bool {
	byteHash := []byte(hash)

	raw, err := argon2.Decode(byteHash)

	if err != nil {
		log.Fatalf("Error during decoding hash: %s\n", err.Error())
	}

	ok, err2 := raw.Verify([]byte(password))

	if err2 != nil {
		log.Fatalf("Error during verifying hash: %s\n", err.Error())
	}

	return ok
}
