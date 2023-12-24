package main

func mulStandart(matrixA, matrixB [][]float64, n int) [][]float64 {
    matrixC := make([][]float64, n)
    for i := 0; i < n; i++ {
        matrixC[i] = make([]float64, n) 
        for j := 0; j < n; j++ { 
			var a float64
            for k := 0; k < n; k++ {
                a += matrixA[i][k] * matrixB[k][j]
            }
			matrixC[i][j] = a
        }
    }
    return matrixC
}
