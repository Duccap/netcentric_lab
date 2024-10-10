package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hamming(DNA1, DNA2 string) (int, error) {
	if len(DNA1) != len(DNA2) {
		return 0, errors.New("DNA sequences must be of equal length")
	}
	hammingDistance := 0
	for i := range DNA1 {
		if DNA1[i] != DNA2[i] {
			hammingDistance++
		}
	}
	return hammingDistance, nil
}

func GenerateRandomDNA(length int) string {
	dnaBases := []rune{'A', 'C', 'G', 'T'}
	rand.Seed(time.Now().UnixNano())
	dna := make([]rune, length)
	for i := range dna {
		dna[i] = dnaBases[rand.Intn(len(dnaBases))]
	}
	return string(dna)
}

func Ex01(length int) {
	for i := 0; i < 1000; i++ {
		DNA1, DNA2 := GenerateRandomDNA(length), GenerateRandomDNA(length)
		if hammingDistance, err := Hamming(DNA1, DNA2); err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Printf("Test %d: Hamming Distance between %s and %s is %d\n", i+1, DNA1, DNA2, hammingDistance)
		}
	}
}
