package main

import "fmt"

type input [][]int

func newTask() input {
	task := make(input, 4)
	//bt,at,wt,tat
	task = [][]int{[]int{10, 0, 0, 0}, []int{5, 0, 0, 0}, []int{8, 0, 0, 0}}

	return task
}

func (in input) fifo() {
	var avwt int
	var avtat int
	in[0][2] = 0
	for i := 1; i < len(in); i++ {
		in[i][2] = 0
		for j := 0; j < i; j++ {
			in[i][2] += in[j][0]
		}
	}

	for i := 0; i < len(in); i++ {
		in[i][3] = in[i][0] + in[i][2]
		avwt += in[i][2]
		avtat += in[i][3]
	}

	avwt, avtat = in.avTime(avwt, avtat)

	fmt.Println("Average Waiting Time: ", avwt)
	fmt.Println("Average Turnaround Time: ", avtat)
	fmt.Println(in)
}

func (in input) RoundRobin() {
	var avwt int
	var avtat int
	qt := 2
	proc := len(in)
	tempBT := []int{}
	t := 0
	for i := 0; i < proc; i++ {
		tempBT = append(tempBT, in[i][0])
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
					in[i][2] = t - in[i][0]
					tempBT[i] = 0
				}
			}
		}

		if done == true {
			break
		}
	}
	for i := 0; i < len(in); i++ {
		in[i][3] = in[i][0] + in[i][2]
		avwt += in[i][2]
		avtat += in[i][3]
	}

	avwt, avtat = in.avTime(avwt, avtat)

	fmt.Println("Average Waiting Time: ", avwt)
	fmt.Println("Average Turnaround Time: ", avtat)
	fmt.Println(in)
}

func (in input) avTime(avwt int, avtat int) (int, int) {
	avwt /= len(in)
	avtat /= len(in)

	return avwt, avtat
}
