package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(3)

	in := newTaskFromFile("task.txt")
	fmt.Println(in)
	quicksort(in)

	schedulingProcess(in)

}
