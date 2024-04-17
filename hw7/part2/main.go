package main

import (
	"fmt"
	"math/rand"
)

const randomNumbersCount = 7

func generateRandomNumbersAndPrintMinMax(start, stop int, randomNumbersCh chan int, minNumberCh, maxNumberCh <-chan int, done chan struct{}) {
	for i := 0; i < randomNumbersCount; i++ {
		randomNumber := rand.Intn(stop) + start
		fmt.Printf("Random number %v is generated\n", randomNumber)
		randomNumbersCh <- randomNumber
	}
	close(randomNumbersCh)

	for min := range minNumberCh {
		fmt.Printf("The min number is %v\n", min)
	}

	for max := range maxNumberCh {
		fmt.Printf("The max number is %v\n", max)
	}

	done <- struct{}{}
	close(done)
}

func findMinMaxNumbers(randomNumbersCh <-chan int, minNumberCh, maxNumberCh chan int) {
	var min, max *int

	for i := range randomNumbersCh {
		if min == nil {
			min = &i
		} else if i < *min {
			min = &i
		}

		if max == nil {
			max = &i
		} else if i > *max {
			max = &i
		}
	}

	minNumberCh <- *min
	close(minNumberCh)

	maxNumberCh <- *max
	close(maxNumberCh)
}

func main() {
	randomNumberCh := make(chan int)
	minNumberCh := make(chan int)
	maxNumberCh := make(chan int)
	done := make(chan struct{})

	go generateRandomNumbersAndPrintMinMax(1, 10, randomNumberCh, minNumberCh, maxNumberCh, done)
	go findMinMaxNumbers(randomNumberCh, minNumberCh, maxNumberCh)

	<-done
}
