package main

import (
	"fmt"
	"math/rand"
)

type input [][]int

func newTask() input {
	task := make(input, 8)
	//bt,at,wt,tat
	//id[0], prior[1], long process[2], bt[3], i/o[4], at[5]
	task = [][]int{[]int{1, 1, 40, 20, 5, 7},
		[]int{2, 1, 60, 10, 10, 0},
		[]int{3, 3, 20, 5, 2, 5},
		[]int{4, 2, 30, 15, 5, 5},
		[]int{5, 3, 20, 10, 5, 5},
		[]int{6, 3, 20, 50, 3, 5},
		[]int{7, 2, 50, 10, 5, 5},
		[]int{8, 2, 10, 30, 3, 5},
		[]int{9, 1, 40, 20, 5, 500}}

	for i := 0; i < len(task); i++ {
		// Sum long process, bt, and i/o
		task[i] = append(task[i], task[i][2]+task[i][3]+task[i][4])
		// add wt and tat
		task[i] = append(task[i], 0)
		task[i] = append(task[i], 0)
		// Removing Long process, bt, and i/o
		task[i] = append(task[i][:2], task[i][5:]...)
	}

	// after the process the slice have id[0], prior[1], at[2], bt[3], wt[4], tat[5]
	return task
}

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
func RoundRobin(in input) [][]int {

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

// SJF NON-PREEMPTIVE ALGORITHM
func SJF(in input) [][]int {

	var avwt int
	var avtat int
	proc := len(in)
	queue := []int{}

	for i := 0; i < proc; i++ {
		pos := i
		for j := i + 1; j < proc; j++ {
			if in[j][3] < in[pos][3] {
				pos = j
			}
		}
		in[i], in[pos] = in[pos], in[i]
	}

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

func avTime(avwt int, avtat int, n int) {
	avwt /= n
	avtat /= n

	fmt.Println("Average Waiting Time: ", avwt)
	fmt.Println("Average Turnaround Time: ", avtat)
}
