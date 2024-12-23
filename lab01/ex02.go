package main

import (
	"fmt"
	"strings"
)

func wordScore(letter string) int {
	switch letter {
	case "A", "E", "I", "O", "U", "L", "N", "R", "S", "T":
		return 1
	case "D", "G":
		return 2
	case "B", "C", "M", "P":
		return 3
	case "F", "H", "V", "W", "Y":
		return 4
	case "K":
		return 5
	case "J", "X":
		return 8
	case "Q", "Z":
		return 10
	default:
		return 0
	}
}


func calculateWord(word string) int {
	total := 0
	for _, letter := range word {
		total += wordScore(strings.ToUpper(string(letter)))
	}
	return total
}

func main() {
	ex := "good day"
	fmt.Println(ex)
	fmt.Println("Total score = ", calculateWord(ex))
}
