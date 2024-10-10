package main

import (
	"fmt"
)

func checkBrackets(s string) bool {
	stack := []rune{}
	matchingBrackets := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, char := range s {
		if openBracket, exists := matchingBrackets[char]; exists {
			if len(stack) > 0 && stack[len(stack)-1] == openBracket {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else if char == '(' || char == '[' || char == '{' {
			stack = append(stack, char)
		}
	}

	return len(stack) == 0
}

func main() {
	example1 := "fmt.Println(a.TypeOf(xyz)){[ ]}"
	example2 := "( [ { } ] )"
	example3 := "[ { ( ] ) }"

	fmt.Println("Example 1:", checkBrackets(example1)) // true
	fmt.Println("Example 2:", checkBrackets(example2)) // true
	fmt.Println("Example 3:", checkBrackets(example3)) // false
	example4 := "{[()()]}"
	example5 := "{[(])}"
	example6 := "{{[[(())]]}}"
	example7 := "{[()]}}"
	example8 := "{[()]}{[()]}"

	fmt.Println("Example 4:", checkBrackets(example4)) // true
	fmt.Println("Example 5:", checkBrackets(example5)) // false
	fmt.Println("Example 6:", checkBrackets(example6)) // true
	fmt.Println("Example 7:", checkBrackets(example7)) // false
	fmt.Println("Example 8:", checkBrackets(example8)) // true
}
