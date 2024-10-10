package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func validateLuhn(numbers string) string {
	trimmedNumbers := strings.Join(regexp.MustCompile(`\s+`).Split(strings.TrimSpace(numbers), -1), "")
	if len(trimmedNumbers) <= 1 {
		return "invalid length"
	}

	sum, isSecondNum := 0, false
	for i := len(trimmedNumbers) - 1; i >= 0; i-- {
		number, err := strconv.Atoi(string(trimmedNumbers[i]))
		if err != nil {
			fmt.Println("Error during conversion")
			return "wrong input"
		}
		if isSecondNum {
			number *= 2
			if number > 9 {
				number -= 9
			}
		}
		sum += number
		isSecondNum = !isSecondNum
	}

	if sum%10 == 0 {
		return "valid"
	}
	return "invalid"
}

func main() {
	fmt.Println("4539    3195     0343     6467: ", validateLuhn("4539 3195 0343 6467"))
	fmt.Println("3: ", validateLuhn("3"))
	fmt.Println("8273 1232 7352 0569: ", validateLuhn("8273 1232 7352 0569"))
}