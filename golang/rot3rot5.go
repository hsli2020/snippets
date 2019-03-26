package main

import (
	"fmt"
	"unicode"
)

// rot13(alphabets) + rot5(numeric)
func rot13rot5(input string) string {

	var result []rune
	rot5map := map[rune]rune{'0': '5', '1': '6', '2': '7', '3': '8', '4': '9', '5': '0', '6': '1', '7': '2', '8': '3', '9': '4'}

	for _, i := range input {
		switch {
		case !unicode.IsLetter(i) && !unicode.IsNumber(i):
			result = append(result, i)
		case i >= 'A' && i <= 'Z':
			result = append(result, 'A'+(i-'A'+13)%26)
		case i >= 'a' && i <= 'z':
			result = append(result, 'a'+(i-'a'+13)%26)
		case i >= '0' && i <= '9':
			result = append(result, rot5map[i])
		case unicode.IsSpace(i):
			result = append(result, ' ')
		}
	}
	return fmt.Sprintf(string(result[:]))
}

func main() {
	text := "ROT18 = ROT13+ROT5. The ROT13 (Caesar cipher by 13 chars) is often combined with ROT5. ROT13 to handle alphabets and ROT5 to handle digits."
	fmt.Println(text)
	fmt.Println(rot13rot5(text))

	fmt.Println("Invertible test:")
	fmt.Println(rot13rot5(rot13rot5(text)))
}
