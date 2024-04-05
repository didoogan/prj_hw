package main

import (
	"fmt"
	"time"
)

func printRoutine(ch chan int) {
	for i := range ch {
		fmt.Println(i)
	}
}

func writeRoutine(readCh, writeCh chan int) {
	for i := range readCh {
		fmt.Println(i)
		writeCh <- i + 1
	}
}

func main() {
	intChan := make(chan int)
	intChan2 := make(chan int)

	intChan <- 5
	intChan <- 6
	intChan <- 8

	go writeRoutine()

	time.Sleep(time.Second)
	//
	go printRutine(intChan2)
	time.Sleep(time.Second)
}
