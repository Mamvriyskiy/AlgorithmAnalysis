package test

import (
	"testing"
	"../sorted"
)

//Слайс элементов отсортированных по убыванию
func BenchmarkHeapSortTestSortedDescending(b *testing.B) {
	sliceA := []int{5, 4, 3, 2, 1}
	sliceB := []int{1, 2, 3, 4, 5}

	sorted.HeapSort(sliceA)

	if !checkSlice(sliceA, sliceB){
		b.Error("Expected", sliceB, "\ngot", sliceA)
	}
}

//Слайс элементов отсортированных по возрастанию
func BenchmarkHeapSortTestSortedAscending(b *testing.B) {
	sliceA := []int{1, 2, 3, 4, 5}
	sliceB := []int{1, 2, 3, 4, 5}

	sorted.HeapSort(sliceA)

	if !checkSlice(sliceA, sliceB) {
		b.Error("Expected", sliceB, "\ngot", sliceA)
	}
}

//Слайс элементов, расположенных в рандомном порядке
func BenchmarkHeapSortTestSortedInconsistency(b *testing.B) {
	sliceA := []int{3, 2, -5, 0, 1}
	sliceB := []int{-5, 0, 1, 2, 3}

	sorted.HeapSort(sliceA)
	
	if !checkSlice(sliceA, sliceB) {
		b.Error("Expected", sliceB, "\ngot", sliceA)
	}
}

//Слайс длиной один
func BenchmarkHeapSortTestSortedLenOne(b *testing.B) {
	sliceA := []int{1}
	sliceB := []int{1}

	sorted.HeapSort(sliceA)
	
	if !checkSlice(sliceA, sliceB) {
		b.Error("Expected", sliceB, "\ngot", sliceA)
	}
}

//Пустой слайс
func BenchmarkHeapSortTestSortedEmptySlice(b *testing.B) {
	sliceA := []int{}
	sliceB := []int{}

	sorted.HeapSort(sliceA)
	
	if !checkSlice(sliceA, sliceB) {
		b.Error("Expected", sliceB, "\ngot", sliceA)
	}
}

