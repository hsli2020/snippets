package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"

	"golang.org/x/crypto/scrypt"
)

func aesKeyFromPassword(password string) ([]byte, error) {
	// DO NOT use this salt value; generate your own random salt. 8 bytes is
	// a good length. Keep the salt secret.
	secretSalt := []byte{0xbc, 0x1e, 0x07, 0xd7, 0xb2, 0xa2, 0x5e, 0x2c}
	return scrypt.Key([]byte(password), secretSalt, 32768, 8, 1, 32)
}

func aesGcmEncrypt(unencrypted []byte, password string) ([]byte, error) {
	key, err := aesKeyFromPassword(password)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// generate a random nonce (makes encryption stronger)
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	encrypted := gcm.Seal(nil, nonce, unencrypted, nil)
	// we need nonce for decryption so we put it at the beginning
	// of encrypted text
	return append(nonce, encrypted...), nil
}

func aesGcmDecrypt(encrypted []byte, password string) ([]byte, error) {
	key, err := aesKeyFromPassword(password)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	if len(encrypted) < gcm.NonceSize() {
		return nil, errors.New("Invalid data")
	}

	// extract random nonce we added to the beginning of the file
	nonce := encrypted[:gcm.NonceSize()]
	encrypted = encrypted[gcm.NonceSize():]

	return gcm.Open(nil, nonce, encrypted, nil)
}

func main() {
	password := "my password"
	d, err := ioutil.ReadFile("main.go")
	if err != nil {
		log.Fatalf("ioutil.ReadFile() failed with %s\n", err)
	}
	encrypted, err := aesGcmEncrypt(d, password)
	if err != nil {
		log.Fatalf("aesGcmEncrypt() failed with %s\n", err)
	}
	decrypted, err := aesGcmDecrypt(encrypted, password)
	if err != nil {
		log.Fatalf("aesGcmDecrypt() failed with %s\n", err)
	}
	if !bytes.Equal(d, decrypted) {
		log.Fatalf("decryption created data different than original\n")
	} else {
		fmt.Printf("Encryption in decryption worked!\n")
	}
}