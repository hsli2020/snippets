package main

import (
	"crypto/rand"
	"fmt"
)

func generateToken() string {
	b := make([]byte, 12)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func getToken(length int) string {
    randomBytes := make([]byte, 32) // length*2
    _, err := rand.Read(randomBytes)
    if err != nil {
        panic(err)
    }
    return base32.StdEncoding.EncodeToString(randomBytes)[:length]
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(generateToken())
		fmt.Println(getToken(24))
	}
}