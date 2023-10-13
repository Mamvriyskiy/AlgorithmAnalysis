package algorithms

import (
	//"fmt"
)

func MatrixLevenshtein(s1, s2 []rune) int {
	if len(s1) == 0 || len(s2) == 0 {
		return 0;
	}

	lenS1 := len(s1) + 1
	lenS2 := len(s2) + 1

	matrix := createMatrix(lenS2, lenS1)

	initMatrix(matrix, lenS2, lenS1)

	return LM(matrix, lenS2, lenS1, s1, s2)
}

func LM(matrix [][]int, a, b int, s1, s2 []rune) int { 
	for i := 1; i < a; i++ {
		r2 := s2[i - 1]
		for j := 1; j < b; j++ {
			r1 := s1[j - 1]
			if r1 != r2 {
				matrix[i][j] = min(matrix[i][j - 1] + 1, matrix[i - 1][j] + 1, matrix[i - 1][j - 1] + 1)
			} else {
				matrix[i][j] = min(matrix[i][j - 1] + 1, matrix[i - 1][j] + 1, matrix[i - 1][j - 1])
			}
		}
	}

	return matrix[a - 1][b - 1]
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
