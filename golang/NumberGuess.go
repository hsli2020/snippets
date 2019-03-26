package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var (
	userInput string
	attempts  int = 0
)

// https://www.socketloop.com/tutorials/golang-random-integer-with-rand-seed-within-a-given-range
func numRandom(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func main() {
	numberToGuess := numRandom(1, 10)
	//fmt.Println(numberToGuess)

	for {
		attempts++
		fmt.Println("A number is generated between 1 to 10. What do you think the number is: ")

		// see https://www.socketloop.com/tutorials/golang-fix-fmt-scanf-on-windows-will-scan-input-twice-problem
		// for scanf under Windows OS
		fmt.Scanf("%v\n", &userInput)

		// check if userinput is integer
		userNumber, err := strconv.ParseInt(userInput, 10, 0)
		if err != nil {
			fmt.Println("You must enter an integer value.")
		} else if numberToGuess == int(userNumber) {
			break
		} else if int(userNumber) < numberToGuess {
			fmt.Printf("Smaller than the number. Try again\n")
		} else {
			fmt.Printf("Larger than the number. Try again\n")
		}
	}
	fmt.Printf("Yup, the number is %v. You guessed it correctly after %v attempts\n", numberToGuess, attempts)
}
