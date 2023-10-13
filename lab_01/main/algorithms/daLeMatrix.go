package algorithms

func MatrixDamerauLevenshtein(s1, s2 []rune) int {
	if len(s1) == 0 || len(s2) == 0 {
		return 0;
	}

	lenS1 := len(s1) + 1
	lenS2 := len(s2) + 1

	matrix := createMatrix(lenS2, lenS1)

	initMatrix(matrix, lenS2, lenS1)

	dist := DLM(matrix, lenS2, lenS1, s1, s2)

	return dist
}

func DLM(matrix [][]int, a, b int, s1, s2 []rune) int {
	for i := 1; i < a; i++ {
		for j := 1; j < b; j++ {
			if s1[j - 1] != s2[i - 1] {
				matrix[i][j] = min(matrix[i][j - 1] + 1, matrix[i - 1][j] + 1, matrix[i - 1][j - 1] + 1)
			} else {
				matrix[i][j] = min(matrix[i][j - 1] + 1, matrix[i - 1][j] + 1, matrix[i - 1][j - 1])
			}

			if i > 2 && j > 2 && s1[j - 1] == s2[i - 2] && s1[j - 2] == s2[i - 1] {
				matrix[i][j] = min(matrix[i][j], matrix[i - 2][j - 2] + 1)
			}
		}
	}

	return matrix[a - 1][b - 1]
}

func initINFMatrix(matrix [][]int, a, b int) {
	for i := 1; i < a; i++ {
		for j := 1; j < b; j++ {
			matrix[i][j] = -1
		}
	}
}


func initMatrixEx(matrix [][]int, a, b int) {
	initMatrix(matrix, a, b)
	initINFMatrix(matrix, a, b)
}

func RecursiveDL(s1, s2 []rune) int {
	return DL(s1, s2)
}

func DL(s1, s2 []rune) int {
	lenB := len(s1)
	lenA := len(s2)

	dist := 0 

	if lenA == 0 {
		dist = lenB
	} else if lenB == 0 {
		dist = lenA
	} else {
		if s1[lenB - 1] != s2[lenA - 1] {
			dist = min(DL(s1, s2[ : lenA - 1]) + 1, DL(s1[ : lenB - 1], s2) + 1, DL(s1[ : lenB - 1], s2[ : lenA - 1]) + 1)
		} else {
			dist = min(DL(s1, s2[ : lenA - 1]) + 1, DL(s1[ : lenB - 1], s2) + 1, DL(s1[ : lenB - 1], s2[ : lenA - 1]))
		}

		if lenB > 1 && lenA > 1 && s1[lenB - 1] == s2[lenA - 2] && s1[lenB - 2] == s2[lenA - 1] {
			dist = min(dist, DL(s1[ : lenB - 2], s2[ : lenA - 2]) + 1)
		}
	}

	return dist
}

func RecursiveDLCash(s1, s2 []rune) int {
	lenS1 := len(s1) + 1
	lenS2 := len(s2) + 1

	matrix := createMatrix(lenS2, lenS1)

	initMatrixEx(matrix, lenS2, lenS1)

	return DLCash(matrix, s1, s2)
}

func DLCash(matrix [][]int, s1, s2 []rune) int {
	lenB := len(s1)
	lenA := len(s2)

	dist := 0

	if matrix[lenA][lenB] == -1 {
		if s1[lenB - 1] != s2[lenA - 1] {
			dist = min(DLCash(matrix, s1, s2[ : lenA - 1]) + 1, DLCash(matrix, s1[ : lenB - 1], s2) + 1, DLCash(matrix, s1[ : lenB - 1], s2[ : lenA - 1]) + 1)
		} else {
			dist = min(DLCash(matrix, s1, s2[ : lenA - 1]) + 1, DLCash(matrix, s1[ : lenB - 1], s2) + 1, DLCash(matrix, s1[ : lenB - 1], s2[ : lenA - 1]))
		}

		if lenB > 1 && lenA > 1 && s1[lenB - 1] == s2[lenA - 2] && s1[lenB - 2] == s2[lenA - 1] {
			dist = min(dist, DLCash(matrix, s1[ : lenB - 2], s2[ : lenA - 2]) + 1)
		}

		matrix[lenA][lenB] = dist
	} else {
		dist = matrix[lenA][lenB]
	}

	return dist
}

