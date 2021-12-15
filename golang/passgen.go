package main

import (
	"os"
	"fmt"
	"flag"
	"strings"
	"math/big"
	"crypto/rand"
	"github.com/atotto/clipboard"
)

var (
	length, charsLength int
	useUpper, useLow, useSpecial, useNumbers, outToClipboard bool
)

const (
	charsLow = "abcdefghijklmnopqrstuvwxyz"
	charsUpper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	charsSpecial = "!#$%&*+-./:=?@^_"
	numbers = "0123456789"
)

func init() {
	flag.IntVar(&length, "c", 8, "usage: -c 16 (to set password length to 16)")
	flag.BoolVar(&useLow, "l", false, "usage: -l (to set use lowercase characters)")
	flag.BoolVar(&useUpper, "u", false, "usage: -u (to set use uppercase characters)")
	flag.BoolVar(&useSpecial, "s", false, "usage: -s (to set use special characters)")
	flag.BoolVar(&useNumbers, "n", false, "usage: -n (to set use numbers)")
	flag.BoolVar(&outToClipboard, "clip", false, "usage: -clip (to send password to clipboard)")
}

func main() {
	flag.Parse()

	var chars string

	if !useLow && !useUpper && !useSpecial && !useNumbers {
		charsLength = 78
		chars += charsLow
		chars += charsUpper
		chars += charsSpecial
		chars += numbers
	}
	if useLow {
		chars += charsLow
		charsLength += 26
	}
	if useUpper {
		chars += charsUpper
		charsLength += 26
	}
	if useSpecial {
		chars += charsSpecial
		charsLength += 16
	}
	if useNumbers {
		chars += numbers
		charsLength += 10
	}

	password := make([]string, length)
	charsArr := strings.Split(chars, "")

	for i := 0; i < length; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(charsLength)))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		password[i] = charsArr[n.Int64()]
	}
	
	if outToClipboard {
		clipboard.WriteAll(strings.Join(password[:], ""))
	} else {
		fmt.Println("Your password is: "+strings.Join(password[:], ""))
	}
}