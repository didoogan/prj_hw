package main

import (
	"fmt"
)

func printRoutine(ch chan int) {
	for i := range ch {
		fmt.Println(i + 2)
	}
	close(ch)
}

func writeRoutine(ch chan int, notify chan struct{}) {
	for i := 1; i < 5; i++ {
		ch <- i
	}
	notify <- struct{}{}

}

func main() {
	ch := make(chan int)

	notify := make(chan struct{})

	go writeRoutine(ch, notify)

	go printRoutine(ch)

	<-notify
}
