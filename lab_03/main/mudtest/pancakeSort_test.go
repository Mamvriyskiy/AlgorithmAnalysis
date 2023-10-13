package test

import (
	"fmt"
	"testing"
	"../sorted"
)

func checkSlice(sliceA, sliceB []int) bool {
	if len(sliceA) != len(sliceB) {
		return false
	}

	var err bool = true

	for i := 0; i < len(sliceA) && err; i++ {
		if sliceA[i] != sliceB[i] {
			err = false
		}
	}

	return err
}

//Слайс элементов отсортированных по убыванию
func PancakeSortTestSortedDescending(b *testing.B) {
	sliceA := []int{100, 43, 4, 1, -3, -43, -432, -941}
	sliceB := []int{-943, -432, -43, -3, 1, 4, 43, 100}

	sorted.PancakeSort(sliceA)

	if !checkSlice(sliceA, sliceB){
		b.Error("Expected", sliceB, "\ngot", sliceA)
	}
}

//Слайс элементов отсортированных по возрастанию
func PancakeSortTestSortedAscending(b *testing.B) {
	sliceA := []int{-943, -432, -43, -3, 1, 4, 43, 100}
	sliceB := []int{-943, -432, -43, -3, 1, 4, 43, 100}

	sorted.PancakeSort(sliceA)

	if !checkSlice(sliceA, sliceB) {
		b.Error("Expected", sliceB, "\ngot", sliceA)
	}
}

//Слайс элементов, расположенных в рандомном порядке
func PancakeSortTestSortedInconsistency(b *testing.B) {
	sliceA := []int{100, -3, 1, -43, 43, 4, -432, -943}
	sliceB := []int{-943, -432, -43, -3, 1, 4, 43, 100}

	sorted.PancakeSort(sliceA)
	
	if !checkSlice(sliceA, sliceB) {
		b.Error("Expected", sliceB, "\ngot", sliceA)
	}
}
