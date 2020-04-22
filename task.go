package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
)

type input [][]int

func newTaskFromFile(filename string) (input, input) {
	t := []string{}
	temporaryTask := [][]string{}
	bs, err := os.Open(filename)

	if err != nil {
		// Option 1 : Log the error & return a call to newDeck()
		// Option 2 : Log the error & Quit the program
		fmt.Println("Error : ", err)
		os.Exit(1)
	}

	var s []string
	scanner := bufio.NewScanner(bs)
	for scanner.Scan() {
		s = append(s, scanner.Text())
	}

	for i := 0; i < len(s); i++ {
		t = strings.Split(s[i], " ")
		temporaryTask = append(temporaryTask, t)
	}

	task := make([][]int, len(temporaryTask))

	for i := 0; i < len(temporaryTask); i++ {
		for j := 0; j < len(temporaryTask[0]); j++ {
			n, _ := strconv.Atoi(temporaryTask[i][j])
			task[i] = append(task[i], n)
		}
	}

	initial := make([][]int, len(task))
	copy(initial, task)
	fmt.Println(initial)

	for i := 0; i < len(task); i++ {
		// Sum long process, bt, and i/o
		task[i] = append(task[i], task[i][2]+task[i][3]+task[i][4])
		// add wt and tat
		task[i] = append(task[i], 0)
		task[i] = append(task[i], 0)
		task[i] = append(task[i], 0)
		// Removing Long process, bt, and i/o
		task[i] = append(task[i][:2], task[i][5:]...)
	}

	// after the process the slice have id[0], prior[1], at[2], bt[3], wt[4], tat[5], process[6]
	return task, initial
}

func schedulingProcess(in [][]int) ([][]int, float64, float64) {
	var wg sync.WaitGroup
	containerFifo := [][]int{}
	containerRoundRobin := [][]int{}
	containerShortestJobFirst := [][]int{}

	queueFIFO := [][]int{}
	queueRoundRobin := [][]int{}
	queueShortestJobFirst := [][]int{}

	temporaryFIFO := [][]int{}
	cycle := 20

	clockFIFO := 0
	clockSJF := 0

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
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		for {
			for i := 0; i < len(containerFifo); i++ {
				queueFIFO = append(queueFIFO, containerFifo[0])
				clockFIFO += containerFifo[0][3]
				temporaryFIFO = append(temporaryFIFO, containerFifo[0])

				for j := 1; j < len(containerFifo); j++ {
					if len(temporaryFIFO) >= 7 && clockFIFO >= containerFifo[j][2] {
						processFIFO := make([]int, len(containerFifo[0]))
						copy(processFIFO, containerFifo[j])
						processFIFO[6] = 2
						containerRoundRobin = append(containerRoundRobin, processFIFO)
						containerFifo = append(containerFifo[0:j], containerFifo[j+1:]...)
						// j is decrement for container
						j--
					} else if clockFIFO >= containerFifo[j][2] {
						temporaryFIFO = append(temporaryFIFO, containerFifo[j])
					}
				}
				temporaryFIFO = nil
				containerFifo = containerFifo[1:]
			}

			for i := 0; i < len(containerRoundRobin); i++ {

				if containerRoundRobin[0][3] > cycle {
					temporaryRR := make([]int, len(containerRoundRobin[0]))
					copy(temporaryRR, containerRoundRobin[0])
					temporaryRR[3] -= 20
					temporaryRR[6] = 3
					containerShortestJobFirst = append(containerShortestJobFirst, temporaryRR)
					queueRoundRobin = append(queueRoundRobin, containerRoundRobin[0])
				} else {
					queueRoundRobin = append(queueRoundRobin, containerRoundRobin[0])
				}
				for _, q := range queueRoundRobin {
					if q[3] > cycle {
						q[3] = 20
					}
				}
				containerRoundRobin = containerRoundRobin[1:]
			}

			for i := 0; i < len(containerShortestJobFirst); i++ {
				queueShortestJobFirst = append(queueShortestJobFirst, containerShortestJobFirst[0])
				clockSJF += containerShortestJobFirst[0][3]

				for m := 1; m < len(containerShortestJobFirst); m++ {
					if clockSJF > (containerShortestJobFirst[m][2] + 20) {
						processSJF := make([]int, len(containerShortestJobFirst[0]))
						copy(processSJF, containerShortestJobFirst[m])
						processSJF[6] = 1
						containerFifo = append(containerFifo, processSJF)
						containerShortestJobFirst = append(containerShortestJobFirst[:m], containerShortestJobFirst[m+1:]...)
						m--
					}
				}
				containerShortestJobFirst = containerShortestJobFirst[1:]
			}

			if len(containerFifo) < 1 && len(containerRoundRobin) < 1 && len(containerShortestJobFirst) < 1 {
				defer wg.Done()
				break
			}
		}
	}(&wg)

	wg.Wait()

	fifo(queueFIFO)
	roundRobin(queueRoundRobin)
	shortestJobFirst(queueShortestJobFirst)

	var n int

	result := [][]int{}
	if len(queueFIFO) >= len(queueRoundRobin) && len(queueFIFO) >= len(queueShortestJobFirst) {
		n = len(queueFIFO)
	} else if len(queueRoundRobin) >= len(queueShortestJobFirst) {
		n = len(queueRoundRobin)
	} else {
		n = len(queueShortestJobFirst)
	}

	for i := 0; i < n; i++ {
		if i < len(queueFIFO) {
			result = append(result, queueFIFO[i])
		}
		if i < len(queueRoundRobin) {
			result = append(result, queueRoundRobin[i])
		}
		if i < len(queueShortestJobFirst) {
			result = append(result, queueShortestJobFirst[i])
		}
	}

	avwt, avtat := avTime(result)

	return result, avwt, avtat
}

//implement fifo algorithm
func fifo(in input) [][]int {

	var avwt int
	var avtat int
	// n := len(in)
	in[0][4] = 0
	for i := 1; i < len(in); i++ {
		in[i][4] = 0
		for j := 0; j < i; j++ {
			in[i][4] += in[j][3]
		}
	}

	for i := 0; i < len(in); i++ {
		in[i][5] = in[i][3] + in[i][4]
		avwt += in[i][4]
		avtat += in[i][5]
	}

	// go avTime(avwt, avtat, n)
	return in
}

// RoundRobin ALGORITHM
func roundRobin(in input) [][]int {

	var avwt int
	var avtat int
	qt := 5
	proc := len(in)
	tempBT := []int{}
	t := 0
	for i := 0; i < proc; i++ {
		tempBT = append(tempBT, in[i][3])
	}

	for {
		done := true

		for i := 0; i < proc; i++ {
			if tempBT[i] > 0 {
				done = false
				if tempBT[i] > qt {
					t += qt

					tempBT[i] -= qt
				} else {
					t = t + tempBT[i]
					in[i][4] = t - in[i][3]
					tempBT[i] = 0
				}
			}
		}

		if done == true {
			break
		}
	}
	for i := 0; i < len(in); i++ {
		in[i][5] = in[i][3] + in[i][4]
		avwt += in[i][4]
		avtat += in[i][5]
	}

	// in.avTime(avwt, avtat)
	return in
}

// ShortestJobFirst NON-PREEMPTIVE ALGORITHM
func shortestJobFirst(in input) [][]int {

	var avwt int
	var avtat int
	proc := len(in)
	queue := []int{}

	// for i := 0; i < proc; i++ {
	// 	pos := i
	// 	for j := i + 1; j < proc; j++ {
	// 		if in[j][3] < in[pos][3] {
	// 			pos = j
	// 		}
	// 	}
	// 	in[i], in[pos] = in[pos], in[i]
	// }

	for i := 0; i < proc; i++ {
		for j := i + 1; j < proc; j++ {
			if in[i][3] > in[j][2] {
				queue = append(queue, j)
			}
		}
		for k := 0; k < len(queue); k++ {
			for l := k + 1; l < len(queue); l++ {
				if in[queue[k]][3] > in[queue[l]][3] {
					in[queue[k]], in[queue[l]] = in[queue[l]], in[queue[k]]
				}
			}
		}
	}

	in[0][4] = 0

	for i := 1; i < proc; i++ {
		in[0][4] = 0
		for j := 0; j < i; j++ {
			in[i][4] += in[j][3]
		}
		avwt += in[i][4]
	}

	for i := 0; i < proc; i++ {
		in[i][5] = in[i][3] + in[i][4]
		avtat += in[i][5]
	}

	// in.avTime(avwt, avtat)
	return in
}

func quicksort(a [][]int) input {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a)

	a[pivot], a[right] = a[right], a[pivot]

	for i := range a {
		if a[i][2] < a[right][2] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	quicksort(a[:left][:])
	quicksort(a[left+1:][:])

	return a
}

func avTime(in [][]int) (float64, float64) {
	var waitingTime float64
	var turnAroundtime float64

	for _, e := range in {
		waitingTime += float64(e[4])
		turnAroundtime += float64(e[5])
	}
	n := float64(len(in))
	avwt := waitingTime / n
	avtat := turnAroundtime / n

	return avwt, avtat
}
