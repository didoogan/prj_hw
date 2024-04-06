package main

import (
	"fmt"
	"math/rand"
)

const randomNumbersCount = 7

func generateRandomNumbers(randomNumbersCh chan int) {
	for i := 0; i < randomNumbersCount; i++ {
		randomNumber := rand.Intn(10) + 1
		fmt.Printf("Random number %v is generated\n", randomNumber)
		randomNumbersCh <- randomNumber
	}

	close(randomNumbersCh)
}

func findAverageNumber(randomNumbersCh <-chan int, averageNumberCh chan float64) {
	var sum int
	var counter int

	for i := range randomNumbersCh {
		sum += i
		counter++
	}

	averageNumber := float64(sum) / float64(counter)
	averageNumberCh <- averageNumber

	close(averageNumberCh)
}

func printAverageNumber(averageNumberCh <-chan float64, done chan struct{}) {
	for i := range averageNumberCh {
		fmt.Printf("\nThe average number is %v\n", i)
	}

	done <- struct{}{}

	close(done)
}

func main() {
	randomNumbersCh := make(chan int)
	averageNumberCh := make(chan float64)
	done := make(chan struct{})

	go generateRandomNumbers(randomNumbersCh)
	go findAverageNumber(randomNumbersCh, averageNumberCh)
	go printAverageNumber(averageNumberCh, done)

	<-done
}
