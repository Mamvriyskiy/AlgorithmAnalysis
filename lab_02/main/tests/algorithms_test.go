package test

import (
	"testing"
	"../algorithms"
	//"fmt"
)

//простой тест
func Benchmark(b *testing.B) {
    r1 := []rune("лабрадор")
	r2 := []rune("гибралтар")

	//fmt.Println("s1 = лабрадор, s2 = гибралтар")
	dist := algorithms.MatrixLevenshtein(r1, r2)
	
    if dist != 5 {
        b.Error("Expected 5, got ", dist)
    }
}