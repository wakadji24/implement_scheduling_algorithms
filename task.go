package main

import (
	"fmt"
	"math/rand"
)

type input [][]int

func newTask() input {
	task := make(input, 8)
	//bt,at,wt,tat
	//id[0], prior[1], long process[2], bt[3], i/o[4], at[5], wt[6], tat[7]
	task = [][]int{[]int{1, 1, 40, 20, 5, 0, 0, 0}, []int{2, 1, 60, 10, 10, 0, 0, 0}, []int{3, 3, 20, 5, 2, 5, 0, 0}, []int{4, 2, 30, 15, 5, 5, 0, 0}}

	return task
}

func (in input) fifo() {
	var avwt int
	var avtat int
	in[0][6] = 0
	for i := 1; i < len(in); i++ {
		in[i][6] = 0
		for j := 0; j < i; j++ {
			in[i][6] += in[j][3]
		}
	}

	for i := 0; i < len(in); i++ {
		in[i][7] = in[i][3] + in[i][6]
		avwt += in[i][6]
		avtat += in[i][7]
	}

	in.avTime(avwt, avtat)
}

func (in input) RoundRobin() {
	var avwt int
	var avtat int
	qt := 2
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
					in[i][6] = t - in[i][3]
					tempBT[i] = 0
				}
			}
		}

		if done == true {
			break
		}
	}
	for i := 0; i < len(in); i++ {
		in[i][7] = in[i][3] + in[i][6]
		avwt += in[i][6]
		avtat += in[i][7]
	}

	in.avTime(avwt, avtat)
}

func (in input) SJF() {
	var avwt int
	var avtat int
	proc := len(in)
	queue := []int{}

	for i := 0; i < proc; i++ {
		for j := i + 1; j < proc; j++ {
			if in[i][3] > in[j][5] {
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

	in[0][6] = 0

	for i := 1; i < proc; i++ {
		in[0][6] = 0
		for j := 0; j < i; j++ {
			in[i][6] += in[j][3]
		}
		avwt += in[i][6]
	}

	for i := 0; i < proc; i++ {
		in[i][7] = in[i][3] + in[i][6]
		avtat += in[i][7]
	}

	in.avTime(avwt, avtat)
}

func quicksort(a [][]int) input {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a)

	a[pivot], a[right] = a[right], a[pivot]

	for i := range a {
		if a[i][5] < a[right][5] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	quicksort(a[:left][:])
	quicksort(a[left+1:][:])

	return a
}

func (in input) avTime(avwt int, avtat int) {
	avwt /= len(in)
	avtat /= len(in)

	fmt.Println("Average Waiting Time: ", avwt)
	fmt.Println("Average Turnaround Time: ", avtat)
	fmt.Println(in)
}
