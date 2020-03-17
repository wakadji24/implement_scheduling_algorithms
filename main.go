package main

func main() {
	input := newTask()
	input = quicksort(input)
	input.SJF()
	// input.RoundRobin()
}
