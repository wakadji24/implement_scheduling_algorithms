package main

import "fmt"

func main() {
	input := newTask()
	quicksort(input)
	fmt.Println(input)
	// input.RoundRobin()
}
