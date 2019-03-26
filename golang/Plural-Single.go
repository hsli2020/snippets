package main

import (
	"fmt"
	"strings"
)

func Singular(input string) string {
	if !IsCountable(input) {
		return input
	}

	var singularDictionary = map[string]string{
		"are":      "is",
		"analyses": "analysis",
		"alumni":   "alumnus",
		"aliases":  "alias",
		"axes":     "axis",
		//"alumni": "alumnae", // for female - cannot have duplicate in map

		"genii":       "genius",
		"data":        "datum",
		"atlases":     "atlas",
		"appendices":  "appendix",
		"barracks":    "barrack",
		"beefs":       "beef",
		"buses":       "bus",
		"brothers":    "brother",
		"cafes":       "cafe",
		"corpuses":    "corpus",
		"campuses":    "campus",
		"cows":        "cow",
		"crises":      "crisis",
		"ganglions":   "ganglion",
		"genera":      "genus",
		"graffiti":    "graffito",
		"loaves":      "loaf",
		"matrices":    "matrix",
		"monies":      "money",
		"mongooses":   "mongoose",
		"moves":       "move",
		"movies":      "movie",
		"mythoi":      "mythos",
		"lice":        "louse",
		"niches":      "niche",
		"numina":      "numen",
		"octopuses":   "octopus",
		"opuses":      "opus",
		"oxen":        "ox",
		"penises":     "penis",
		"vaginas":     "vagina",
		"vertices":    "vertex",
		"viruses":     "virus",
		"shoes":       "shoe",
		"sexes":       "sex",
		"testes":      "testis",
		"turfs":       "turf",
		"teeth":       "tooth",
		"feet":        "foot",
		"cacti":       "cactus",
		"children":    "child",
		"criteria":    "criterion",
		"news":        "news",
		"deer":        "deer",
		"echoes":      "echo",
		"elves":       "elf",
		"embargoes":   "embargo",
		"foes":        "foe",
		"foci":        "focus",
		"fungi":       "fungus",
		"geese":       "goose",
		"heroes":      "hero",
		"hooves":      "hoof",
		"indices":     "index",
		"knifes":      "knife",
		"leaves":      "leaf",
		"lives":       "life",
		"men":         "man",
		"mice":        "mouse",
		"nuclei":      "nucleus",
		"people":      "person",
		"phenomena":   "phenomenon",
		"potatoes":    "potato",
		"selves":      "self",
		"syllabi":     "syllabus",
		"tomatoes":    "tomato",
		"torpedoes":   "torpedo",
		"vetoes":      "veto",
		"women":       "woman",
		"zeroes":      "zero",
		"natives":     "native",
		"hives":       "hive",
		"quizzes":     "quiz",
		"bases":       "basis",
		"diagnostic":  "diagnosis",
		"parentheses": "parenthesis",
		"prognoses":   "prognosis",
		"synopses":    "synopsis",
		"theses":      "thesis",
	}

	result := singularDictionary[strings.ToLower(input)]

	if result == "" {
		// to handle words like apples, doors, cats
		if len(input) > 2 {
			if string(input[len(input)-1]) == "s" {
				return string(input[:len(input)-1])
			}
		}
		return input
	} else {
		return result
	}
}

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

	fmt.Println("Are :", Singular("Are"))
	fmt.Println("Equipment :", Singular("Equipment"))
	fmt.Println("is :", Singular("is"))
	fmt.Println("Apples :", Singular("Apples"))
	fmt.Println("oranges :", Singular("oranges"))
	fmt.Println("bases :", Singular("bases"))
	fmt.Println("Doors :", Singular("Doors"))

}
