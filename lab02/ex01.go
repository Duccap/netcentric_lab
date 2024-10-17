package main

import (
	"fmt"
	"sync"
)

func countCharacters(input string, wg *sync.WaitGroup, ch chan<- map[rune]int) {
	defer wg.Done()

	charCount := make(map[rune]int)
	for _, char := range input {
		charCount[char]++
	}

	ch <- charCount
}

func main() {
	input := "Success is not final, failure is not fatal: It is the courage to continue that counts."
	numberOfChunks := 12

	ch := make(chan map[rune]int) 
	var wg sync.WaitGroup

	chunkSize := len(input) / numberOfChunks 
	var startIndex, endIndex int
	for i := 0; i < numberOfChunks; i++ {
		startIndex = i * chunkSize
		endIndex = (i + 1) * chunkSize
		if i == numberOfChunks-1 {
			endIndex = len(input)
		}
		wg.Add(1)
		go countCharacters(input[startIndex:endIndex], &wg, ch)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	result := make(map[string]int)
	for charCount := range ch {
		for char, count := range charCount {
			result[string(char)] += count
		}
	}

	fmt.Println("Character counts:", result)
}