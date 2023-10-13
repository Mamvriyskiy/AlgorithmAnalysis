package algorithms

// import "fmt"

//алгоритм Виноградова 
func MulVinogradovA(matrixA, matrixB [][]int, a, b, c int) [][]int {
    row := make([]int, a)
    column := make([]int, c)
    matrixC := make([][]int, a)

    for i := 0; i < a; i++ {
        matrixC[i] = make([]int, c)
        for j := 0; j < b / 2 ; j++ {
            row[i] = row[i] + matrixA[i][2 * j] * matrixA[i][2 * j + 1]
        }
    }

    for i := 0; i < c; i++ {
        for j := 0; j < b / 2; j++ {
            column[i] = column[i] + matrixB[2 * j][i]*matrixB[2 * j + 1][i] 
        }
    }

    for i := 0; i < a; i++ {
        for j := 0; j < c; j++ {
            matrixC[i][j] = -row[i] - column[j]
            for k := 0; k < b / 2; k++ {
                matrixC[i][j] = matrixC[i][j] + (matrixA[i][2 * k + 1] + matrixB[2 * k][j]) *
                    (matrixA[i][2 * k] + matrixB[2 * k + 1][j])
            }
        }
    }

    if (b % 2 == 1) {
        for i := 0; i < a; i++ {
            for j := 0; j < c; j++ {
                matrixC[i][j] = matrixC[i][j] + matrixA[i][b - 1] * matrixB[b - 1][j]
            }
        }
    }   

    return matrixC
}

//оптимизированный алгоритм Виноградова(добвлен буфер, побитовый сдвиг, замена b / 2 на переменную)
func MulVinogradovB(matrixA, matrixB [][]int, a, b, c int) [][]int {
    row := make([]int, a)
    column := make([]int, c)
    matrixC := make([][]int, a)

    num := b / 2

    for i := 0; i < a; i++ {
        matrixC[i] = make([]int, c)
        buf := 0
        for j := 0; j < num ; j++ {
            buf += matrixA[i][j << 1] * matrixA[i][j << 1 + 1]
        }
        row[i] = buf
    }

    for i := 0; i < c; i++ {
        buf := 0
        for j := 0; j < num; j++ {
            buf += matrixB[j << 1][i] * matrixB[j << 1 + 1][i] 
        }
        column[i] = buf
    }

    for i := 0; i < a; i++ {
        for j := 0; j < c; j++ {
            buf := -row[i] - column[j]
            for k := 0; k < num; k++ {
                buf += (matrixA[i][k << 1 + 1] + matrixB[k << 1][j]) *
                    (matrixA[i][k << 1] + matrixB[k << 1 + 1][j])
            }
            matrixC[i][j] = buf
        }
    }

    if (b % 2 == 1) {
        l := b - 1
        for i := 0; i < a; i++ {
            for j := 0; j < c; j++ {
                matrixC[i][j] += matrixA[i][l] * matrixB[l][j]
            }
        }
    }   

    return matrixC
}

//алгоритм простого умножения
func MulMAtrixA(matrixA, matrixB [][]int, n, m, p int) [][]int {
    matrixC := make([][]int, n)
    for i := 0; i < n; i++ {
        matrixC[i] = make([]int, m) 
        for j := 0; j < m; j++ { 
            for k := 0; k < p; k++ {
                matrixC[i][j] = matrixC[i][j] + matrixA[i][k] * matrixB[k][j]
            }
        }
    }

    return matrixC
}

//оптимизированный алгоритм простого умножения
func MulMAtrixB(matrixA, matrixB [][]int, n, m, p int) [][]int {
    matrixC := make([][]int, n)
    for i := 0; i < n; i++ {
        matrixC[i] = make([]int, n) 
        for j := 0; j < m; j++ { 
			var a int
            for k := 0; k < p; k++ {
                a += matrixA[i][k] * matrixB[k][j]
            }
			matrixC[i][j] = a
        }
    }
    return matrixC
}


