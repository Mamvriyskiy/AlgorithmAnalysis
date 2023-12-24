package test

import (
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
func BenchmarkPancakeSortTestSortedDescending(b *testing.B) {
	sliceA := []int{5, 4, 3, 2, 1}
	sliceB := []int{1, 2, 3, 4, 5}

	sorted.PancakeSort(sliceA)

	if !checkSlice(sliceA, sliceB){
		b.Error("Expected", sliceB, "\ngot", sliceA)
	}
}

//Слайс элементов отсортированных по возрастанию
func BenchmarkPancakeSortTestSortedAscending(b *testing.B) {
	sliceA := []int{1, 2, 3, 4, 5}
	sliceB := []int{1, 2, 3, 4, 5}

	sorted.PancakeSort(sliceA)

	if !checkSlice(sliceA, sliceB) {
		b.Error("Expected", sliceB, "\ngot", sliceA)
	}
}

//Слайс элементов, расположенных в рандомном порядке
func BenchmarkPancakeSortTestSortedInconsistency(b *testing.B) {
	sliceA := []int{3, 2, -5, 0, 1}
	sliceB := []int{-5, 0, 1, 2, 3}

	sorted.PancakeSort(sliceA)
	
	if !checkSlice(sliceA, sliceB) {
		b.Error("Expected", sliceB, "\ngot", sliceA)
	}
}

//Слайс длиной один
func BenchmarkPancakeSortTestSortedLenOne(b *testing.B) {
	sliceA := []int{1}
	sliceB := []int{1}

	sorted.PancakeSort(sliceA)
	
	if !checkSlice(sliceA, sliceB) {
		b.Error("Expected", sliceB, "\ngot", sliceA)
	}
}

//Пустой слайс
func BenchmarkPancakeSortTestSortedEmptySlice(b *testing.B) {
	sliceA := []int{}
	sliceB := []int{}

	sorted.PancakeSort(sliceA)
	
	if !checkSlice(sliceA, sliceB) {
		b.Error("Expected", sliceB, "\ngot", sliceA)
	}
}