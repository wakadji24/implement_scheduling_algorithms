package main

import (
	"fmt"
)

func main() {
	queueFIFO := [][]int{}
	queueRR := [][]int{}
	queueSJF := [][]int{}

	fifoalgorithm := [][]int{}

	times := 0
	in := newTask()
	quicksort(in)

	fmt.Println(in)

	for i := 0; i < len(in); i++ {
		if in[i][1] == 1 && len(queueFIFO) < 8 && in[i][2] <= times {
			queueFIFO = append(queueFIFO, in[i])
			times = times + in[i][3]
		} else if in[i][1] == 2 || in[i][1] == 1 {
			queueRR = append(queueRR, in[i])
		} else if in[i][1] == 3 {
			queueSJF = append(queueSJF, in[i])
		}

		if times < in[i][2] || len(queueFIFO) > 7 {
			fifo(queueFIFO)

			for j := 0; j < len(queueFIFO); j++ {
				fifoalgorithm = append(fifoalgorithm, queueFIFO[j])
			}
			queueFIFO = nil
		}
	}

	if queueFIFO != nil {
		fifo(queueFIFO)

		for j := 0; j < len(queueFIFO); j++ {
			fifoalgorithm = append(fifoalgorithm, queueFIFO[j])
		}
		queueFIFO = nil
	}

	fmt.Println(fifoalgorithm)
	fmt.Println(RoundRobin(queueRR))
	fmt.Println(SJF(queueSJF))
	// quicksort(in)
	// in.SJF()
	// fmt.Println(in)
}
