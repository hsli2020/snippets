package main

import (
	"fmt"
	"strings"
)

func rot47(input string) string {

	var result []string
	for i := range input[:len(input)] {
		j := int(input[i])
		if (j >= 33) && (j <= 126) {
			result = append(result, string(rune(33+((j+14)%94))))
		} else {
			result = append(result, string(input[i]))
		}

	}
	return strings.Join(result, "")
}

func main() {

	text := "ROT47 makes text unreadable and it is also invertible algorithm!"
	fmt.Println("ROT47 : ")
	fmt.Println(text)
	fmt.Println(rot47(text))

	// invertible algorithm, apply the same algorithm twice to get the original text
	fmt.Println(rot47(rot47(text)))
}
