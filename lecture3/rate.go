package main

import "fmt"

func rate(s []int) {
	var counter int
	for _, i := range s {
		counter += i
	}

	fmt.Printf("The average reate is: %v\n", float32(counter)/float32(len(s)))
}

func main() {
	r := []int{3, 10, 6, 8, 2, 5, 9, 5, 7}
	rate(r)
}
