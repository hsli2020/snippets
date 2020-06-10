package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

const nonceLen = 20

func generateNonce() (string, error) {
	nonceBytes := make([]byte, nonceLen)
	if _, err := io.ReadAtLeast(rand.Reader, nonceBytes, nonceLen); err != nil {
		return "", fmt.Errorf("io.ReadAtLeast(rand.Reader, nonceBytes, %d): %v", nonceLen, err)
	}
	return base64.StdEncoding.EncodeToString(nonceBytes), nil
}

func main() {
	nonce, _ := generateNonce()
	fmt.Println(nonce)
}
