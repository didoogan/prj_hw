package main

import (
	"fmt"
	"math/rand"
)

func generateRandomNumbers(ch chan int) {
	for i := 0; i < 7; i++ {
		randomNumber := rand.Intn(10) + 1
		fmt.Printf("Random number %v is generated\n", randomNumber)
		ch <- randomNumber
	}

	close(ch)
}

func findAverageNumber(readCh <-chan int, writeCh chan float64) {
	var sum int
	var counter int

	for i := range readCh {
		sum += i
		counter++
	}

	var averageNumber float64
	if counter > 0 {
		averageNumber = float64(sum) / float64(counter)
	}

	writeCh <- averageNumber

	close(writeCh)
}

func printAverageNumber(ch chan float64, done chan struct{}) {
	for i := range ch {
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
