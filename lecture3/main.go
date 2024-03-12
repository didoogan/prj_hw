package main

import "fmt"

func main() {
	var intSlice []int

	for i := 1; i <= 1000; i++ {
		intSlice = append(intSlice, i)
		fmt.Printf("lent = %v, capacity = %v\n", len(intSlice), cap(intSlice))
	}
}
