package main

import (
	"fmt"
	//"math"
)

func matrixLevenshtein(s1, s2 []rune) {
	// lenS1 := utf8.RuneCountInString(s1) + 1
	// lenS2 := utf8.RuneCountInString(s2) + 1

	lenS1 := len(s1) + 1
	lenS2 := len(s2) + 1


	matrix := createMatrix(lenS2, lenS1)

	initMatrix(matrix, lenS2, lenS1)

	createResult(matrix, lenS2, lenS1, s1, s2)

	fmt.Println(matrix[lenS2 - 1][lenS1 - 1])
}

func createResult(matrix [][]int, a, b int, s1, s2 []rune) { 
	for i := 1; i < a; i++ {
		r2 := s2[i - 1]
		for j := 1; j < b; j++ {
			k := 0
			r1 := s1[j - 1]
			//fmt.Println(r1, r2)
			if r1 != r2 {
				k++
			}
			
			matrix[i][j] = min(matrix[i][j - 1] + 1, matrix[i - 1][j] + 1, matrix[i - 1][j - 1] + k)
		}
	}
}

func initMatrix(matrix [][]int, a, b int) {
	for i := 0; i < a; i++ {
		matrix[i][0] = i
	}

	for j := 0; j < b; j++ {
		matrix[0][j] = j
	}
}

func createMatrix(a, b int) [][]int {
	matrix := make([][]int, a)

	for i := 0; i < a; i++ {
		matrix[i] = make([]int, b)
	}

	return matrix
}
