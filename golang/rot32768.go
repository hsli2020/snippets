package main

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

// rot32768 rotates utf8 string
// rot5map rotates digits
func rot32768(input string) string {
	var result []string
	rot5map := map[rune]rune{'0': '5', '1': '6', '2': '7', '3': '8', '4': '9', '5': '0', '6': '1', '7': '2', '8': '3', '9': '4'}

	for _, i := range input {
		switch {
		case unicode.IsSpace(i):
			result = append(result, " ")
		case i >= '0' && i <= '9':
			result = append(result, string(rot5map[i]))
		case utf8.ValidRune(i):
			//result = append(result, string(rune(i) ^ 0x80))
			result = append(result, string(rune(i)^utf8.RuneSelf))
		}

	}

	return strings.Join(result, "")
}

func main() {
	text := "жѳМѭњЂЯёЧВ 一二三 handle alphabets = ขอพฮศโำฐเฦ abc世界你好123 ペツワケザユプルヂザ"
	fmt.Println(text)
	fmt.Println(rot32768(text))

	fmt.Println("Invertible test:")
	fmt.Println(rot32768(rot32768(text)))
}
