package main

func mulVinogradov(matrixA, matrixB [][]float64, n int) [][]float64 {
    row := make([]float64, n)
    column := make([]float64, n)
    matrixC := make([][]float64, n)

    num := n / 2

    for i := 0; i < n; i++ {
        matrixC[i] = make([]float64, n)
        buf := 0.0
        for j := 0; j < num ; j++ {
            buf += matrixA[i][j << 1] * matrixA[i][j << 1 + 1]
        }
        row[i] = buf
    }

    for i := 0; i < n; i++ {
        buf := 0.0
        for j := 0; j < num; j++ {
            buf += matrixB[j << 1][i] * matrixB[j << 1 + 1][i] 
        }
        column[i] = buf
    }

    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            buf := -row[i] - column[j]
            for k := 0; k < num; k++ {
                buf += (matrixA[i][k << 1 + 1] + matrixB[k << 1][j]) *
                    (matrixA[i][k << 1] + matrixB[k << 1 + 1][j])
            }
            matrixC[i][j] = buf
        }
    }

    if (n % 2 == 1) {
        l := n - 1
        for i := 0; i < n; i++ {
            for j := 0; j < n; j++ {
                matrixC[i][j] += matrixA[i][l] * matrixB[l][j]
            }
        }
    }   

    return matrixC
}
