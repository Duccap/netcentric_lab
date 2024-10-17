package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	totalSeats   = 30
	totalStudents = 100
	maxReadingTime = 4
)

type Student struct {
	id        int
	readTime  int
}

var wg sync.WaitGroup
var mu sync.Mutex

func main() {
	rand.Seed(time.Now().UnixNano())
	openingHours := simulateLibraryDay()
	fmt.Printf("Total opening hours: %d\n", openingHours)
}

func simulateLibraryDay() int {
	timeElapsed := 0
	waitingStudents := make(chan Student, totalStudents)
	librarySeats := make(chan Student, totalSeats)

	go generateStudents(waitingStudents)

	// Serve students until all are done
	for servedStudents := 0; servedStudents < totalStudents; {
		// Process students currently seated
		processLibrarySeats(librarySeats, &servedStudents)

		// Admit students if seats are available
		fillSeats(waitingStudents, librarySeats)

		time.Sleep(1 * time.Second) // Simulate 1 hour
		timeElapsed++
	}

	// Wait for all students to leave before closing the library
	wg.Wait()
	return timeElapsed
}

func generateStudents(waitingStudents chan<- Student) {
	for i := 1; i <= totalStudents; i++ {
		readTime := rand.Intn(maxReadingTime) + 1
		student := Student{id: i, readTime: readTime}
		waitingStudents <- student
		fmt.Printf("Time %d: Student %d arrived, wants to read for %d hours\n", 0, student.id, student.readTime)
		time.Sleep(time.Duration(rand.Intn(2)) * time.Millisecond) // Random interval between arrivals
	}
	close(waitingStudents)
}

func fillSeats(waitingStudents <-chan Student, librarySeats chan<- Student) {
	for len(librarySeats) < totalSeats {
		select {
		case student, ok := <-waitingStudents:
			if !ok {
				return // detect if waitingStudents is closed
			}
			librarySeats <- student
			wg.Add(1)
			fmt.Printf("Time %d: Student %d starts reading in the library\n", time.Now().Second(), student.id)
		default:
			return
		}
	}
}

func processLibrarySeats(librarySeats chan Student, servedStudents *int) {
	// Get the current number of students in the library
	currentCount := len(librarySeats)
	for i := 0; i < currentCount; i++ {
		// Take a student from the librarySeats channel
		student := <-librarySeats
		student.readTime--

		if student.readTime == 0 {
			fmt.Printf("Time %d: Student %d is leaving after reading\n", time.Now().Second(), student.id)
			wg.Done()
			*servedStudents++
		} else {
			// If the student still has time left, re-add them to librarySeats
			librarySeats <- student
		}
	}
}

