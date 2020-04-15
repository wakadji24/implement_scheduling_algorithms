package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(3)

	containerFifo := [][]int{}
	containerRoundRobin := [][]int{}
	containerShortestJobFirst := [][]int{}

	queueFIFO := [][]int{}
	queueRoundRobin := [][]int{}
	queueShortestJobFirst := [][]int{}

	in := newTask()
	quicksort(in)

	//fmt.Println(in)

	for i := 0; i < len(in); i++ {
		if in[i][1] == 1 {
			containerFifo = append(containerFifo, in[i])
		}
		if in[i][1] == 2 {
			containerRoundRobin = append(containerRoundRobin, in[i])
		}
		if in[i][1] == 3 {
			containerShortestJobFirst = append(containerShortestJobFirst, in[i])
		}
	}

	queueFIFO, queueRoundRobin, queueShortestJobFirst = schedulingProcess(&wg, containerFifo, containerRoundRobin, containerShortestJobFirst)

	fmt.Println(queueFIFO)
	fmt.Println(queueRoundRobin)
	fmt.Println(queueShortestJobFirst)
	//fmt.Println(queueSJF)
	// quicksort(in)
	// in.SJF()
	// fmt.Println(in)
}
