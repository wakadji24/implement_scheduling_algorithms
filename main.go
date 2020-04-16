package main

import (
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(3)

	in := newTask()
	quicksort(in)

	schedulingProcess(in)

}
