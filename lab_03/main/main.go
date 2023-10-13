package main

import (
	"fmt"
	// "bufio"
	// "os"
	"./sorted"
)

func main() {
	//reader := bufio.NewReader(os.Stdin)

	//var n int

	sliceA := []int{100, 43, 4, 1, -3, -43, -432, -941}

	sorted.PancakeSort(sliceA)

	if false {
		fmt.Println(sliceA)
	}
}
