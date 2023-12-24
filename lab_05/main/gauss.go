package main



func createInverseMatrix(size int) [][]float64 {
	inverse := make([][]float64, size)
	for i := 0; i < size; i++ {
		inverse[i] = make([]float64, size)
		inverse[i][i] = 1
	}
	
	return inverse
}

func gaussMethodEx(matrix [][]float64) [][]float64 {
	size := len(matrix)
	inverse := createInverseMatrix(size)
	gaussMethod(matrix, inverse, size)

	return inverse
}

func swapRows(matrix [][]float64, row1, row2 int) {
	matrix[row1], matrix[row2] = matrix[row2], matrix[row1]
}


func subtractRows(matrix [][]float64, destRow, sourceRow, size int, factor float64) {
	for i := 0; i < size; i++ {
		matrix[destRow][i] -= factor * matrix[sourceRow][i]
	}
}	

func gaussMethod(matrix, inverse [][]float64, size int) {
	for i := 0; i < size; i++ {
		pivot := matrix[i][i]
		if pivot == 0 {
			var k int
			for k = i + 1; i < size && matrix[k][i] == 0; k++ {}
			if k >= size {
				break
			} else {
				swapRows(matrix, i, k)
				swapRows(inverse, i, k)
				pivot = matrix[i][i]
			}
		}

		for j := 0; j < size; j++ {
			if j != i && matrix[j][i] != 0 {
				factor := float64(matrix[j][i]) / float64(pivot)
				subtractRows(matrix, j, i, size, factor)
				subtractRows(inverse, j, i, size, factor)
			}
		}
	}
}
