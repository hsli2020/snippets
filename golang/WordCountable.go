package main

import (
	"fmt"
	"strings"
)

func IsCountable(input string) bool {
	// dictionary of word that has no plural version
	toCheck := strings.ToLower(input)

	var nonCountable = []string{
		"audio",
		"bison",
		"chassis",
		"compensation",
		"coreopsis",
		"data",
		"deer",
		"education",
		"emoji",
		"equipment",
		"fish",
		"furniture",
		"gold",
		"information",
		"knowledge",
		"love",
		"rain",
		"money",
		"moose",
		"nutrition",
		"offspring",
		"plankton",
		"pokemon",
		"police",
		"rice",
		"series",
		"sheep",
		"species",
		"swine",
		"traffic",
		"wheat",
	}

	for _, v := range nonCountable {
		if toCheck == v {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("Traffic is countable? ", IsCountable("traffic"))
	fmt.Println("Swine is countable? ", IsCountable("SwINE"))
	fmt.Println("Fish is countable? ", IsCountable("FISH"))
	fmt.Println("Apple is countable? ", IsCountable("Apple"))
	fmt.Println("People is countable? ", IsCountable("People"))
}
