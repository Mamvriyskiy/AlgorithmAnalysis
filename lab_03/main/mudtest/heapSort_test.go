package test

import (
	"testing"
	"../sorted"
)

//Слайс элементов отсортированных по убыванию
func BenchmarkHeapSortTestSortedDescending(b *testing.B) {
	sliceA := []int{100, 43, 4, 1, -3, -43, -432, -941}
	sliceB := []int{-941, -432, -43, -3, 1, 4, 43, 100}

	sorted.HeapSort(sliceA)

	if !checkSlice(sliceA, sliceB){
		b.Error("Expected", sliceB, "\ngot", sliceA)
	}
}

//Слайс элементов отсортированных по возрастанию
func BenchmarkHeapSortTestSortedAscending(b *testing.B) {
	sliceA := []int{-943, -432, -43, -3, 1, 4, 43, 100}
	sliceB := []int{-943, -432, -43, -3, 1, 4, 43, 100}

	sorted.HeapSort(sliceA)

	if !checkSlice(sliceA, sliceB) {
		b.Error("Expected", sliceB, "\ngot", sliceA)
	}
}

//Слайс элементов, расположенных в рандомном порядке
func BenchmarkHeapSortTestSortedInconsistency(b *testing.B) {
	sliceA := []int{100, -3, 1, -43, 43, 4, -432, -943}
	sliceB := []int{-943, -432, -43, -3, 1, 4, 43, 100}

	sorted.HeapSort(sliceA)
	
	if !checkSlice(sliceA, sliceB) {
		b.Error("Expected", sliceB, "\ngot", sliceA)
	}
}
