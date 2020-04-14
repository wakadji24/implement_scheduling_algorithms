package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(3)

	queueFIFO := [][]int{}
	queueRR := [][]int{}
	queueSJF := [][]int{}

	containerFIFO := [][]int{}
	containerRR := [][]int{}
	containerSJF := [][]int{}

	temporaryFIFO := [][]int{}
	//temporarySJF := [][]int{}
	cycle := 20

	clockFIFO := 0
	clockSJF := 0
	in := newTask()
	quicksort(in)

	//fmt.Println(in)

	for i := 0; i < len(in); i++ {
		if in[i][1] == 1 {
			containerFIFO = append(containerFIFO, in[i])
		}
		if in[i][1] == 2 {
			containerRR = append(containerRR, in[i])
		}
		if in[i][1] == 3 {
			containerSJF = append(containerSJF, in[i])
		}
	}
	fmt.Println("FIFO:", containerFIFO)
	fmt.Println("RR:", containerRR)
	fmt.Println("SJF:", containerSJF)

	wg.Add(3)
	go func() {
		for i := 0; i < len(containerFIFO); i++ {
			queueFIFO = append(queueFIFO, containerFIFO[i])
			clockFIFO += containerFIFO[i][3]
			temporaryFIFO = append(temporaryFIFO, containerFIFO[i])
			for j := i + 1; j < len(containerFIFO); j++ {
				if len(temporaryFIFO) >= 7 && clockFIFO >= containerFIFO[j][2] {
					containerRR = append(containerRR, containerFIFO[j])
					containerFIFO = append(containerFIFO[0:j], containerFIFO[j+1:]...)
					// j is decrement for container
					j--
				} else if clockFIFO >= containerFIFO[j][2] {
					temporaryFIFO = append(temporaryFIFO, containerFIFO[j])
				}
			}
			temporaryFIFO = nil
		}
		defer wg.Done()
	}()

	go func() {

		for k := 0; k < len(containerRR); k++ {
			if containerRR[k][3] > cycle {
				temporaryRR := make([]int, len(containerRR[k]))
				copy(temporaryRR, containerRR[k])
				temporaryRR[3] -= 20
				containerSJF = append(containerSJF, temporaryRR)
				queueRR = append(queueRR, containerRR[k])
			} else {
				queueRR = append(queueRR, containerRR[k])
			}
			for _, q := range queueRR {
				if q[3] > cycle {
					q[3] = 20
				}
			}
		}
		defer wg.Done()
	}()

	go func() {

		for l := 0; l < len(containerSJF); l++ {
			queueSJF = append(queueSJF, containerSJF[l])
			clockSJF += containerSJF[l][3]
			for m := l + 1; m < len(containerSJF); m++ {
				if clockSJF > (containerSJF[m][2] + 20) {
					containerFIFO = append(containerFIFO, containerSJF[m])
					containerSJF = append(containerSJF[:m], containerSJF[m+1:]...)
					m--
				}
			}
		}
		defer wg.Done()
	}()
	wg.Wait()
	fmt.Println("============================================================")
	fmt.Println("FIFO:", queueFIFO)
	fmt.Println("RR:", queueRR)
	fmt.Println("SJF:", queueSJF)
	//fmt.Println(queueSJF)
	// quicksort(in)
	// in.SJF()
	// fmt.Println(in)
}
