package test

import (
	"testing"
	"../sorted"
)

//Слайс элементов отсортированных по убыванию
func BenchmarkShakerSortTestSortedDescending(b *testing.B) {
	sliceA := []int{5, 4, 3, 2, 1}
	sliceB := []int{1, 2, 3, 4, 5}

	sorted.ShakerSort(sliceA)

	if !checkSlice(sliceA, sliceB){
		b.Error("Expected", sliceB, "\ngot", sliceA)
	}
}

//Слайс элементов отсортированных по возрастанию
func BenchmarkShakerSortTestSortedAscending(b *testing.B) {
	sliceA := []int{1, 2, 3, 4, 5}
	sliceB := []int{1, 2, 3, 4, 5}

	sorted.ShakerSort(sliceA)

	if !checkSlice(sliceA, sliceB) {
		b.Error("Expected", sliceB, "\ngot", sliceA)
	}
}

//Слайс элементов, расположенных в рандомном порядке
func BenchmarkShakerSortTestSortedInconsistency(b *testing.B) {
	sliceA := []int{3, 2, -5, 0, 1}
	sliceB := []int{-5, 0, 1, 2, 3}

	sorted.ShakerSort(sliceA)
	
	if !checkSlice(sliceA, sliceB) {
		b.Error("Expected", sliceB, "\ngot", sliceA)
	}
}

//Слайс длиной один
func BenchmarkShakerSortTestSortedLenOne(b *testing.B) {
	sliceA := []int{1}
	sliceB := []int{1}

	sorted.ShakerSort(sliceA)
	
	if !checkSlice(sliceA, sliceB) {
		b.Error("Expected", sliceB, "\ngot", sliceA)
	}
}

//Пустой слайс
func BenchmarkShakerSortTestSortedEmptySlice(b *testing.B) {
	sliceA := []int{}
	sliceB := []int{}

	sorted.ShakerSort(sliceA)
	
	if !checkSlice(sliceA, sliceB) {
		b.Error("Expected", sliceB, "\ngot", sliceA)
	}
}