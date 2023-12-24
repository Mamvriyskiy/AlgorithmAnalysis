package main

import (
	"fmt"
	"math/rand"
)

type matrixStruct struct {
	matrixA [][]float64
	matrixB [][]float64
	inverseMatrix [][]float64
	standartMul [][]float64
	vinogradMul [][]float64
	size int
	uniq int
}

func readMatrix(size int) (error, [][]float64) {
	var err error
	matrix := make([][]float64, size)
	for i := 0; i < size; i++ {
		matrix[i] = make([]float64, size)
		for j := 0; j < size; j++ {
			fmt.Printf("Введите %d элемент %d строки: ", j + 1, i + 1)
			_, err = fmt.Scanf("%f", &matrix[i][j])
			if err != nil {
				return err, matrix
			}
		}
	}

	return err, matrix
}

func createRandomMatrix(n int) [][]float64 {
	mtr := make([][]float64, n)
	for i := 0; i < n; i++ {
		mtr[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			mtr[i][j] = rand.Float64() * 100
		}
	}

	mtr[n - 1][n - 1] = 0
 	return mtr
}

func createMatrixEx(count, size, flag int) (error, []matrixStruct) {
	matrixSlice := make([]matrixStruct, count)
	var err error
	for i := 0; i < count && err == nil; i++ {
		if flag == 1 {
			fmt.Println()
			fmt.Println("Ввод матрицы A")
			err, matrixSlice[i].matrixA = readMatrix(size)
			fmt.Println()
			fmt.Println("Ввод матрицы B")
			err, matrixSlice[i].matrixB = readMatrix(size)
		} else {
			matrixSlice[i].matrixA = createRandomMatrix(size)
			matrixSlice[i].matrixB = createRandomMatrix(size)
		}
		matrixSlice[i].size = size
		matrixSlice[i].uniq = i + 1
	}

	return err, matrixSlice
}

func createSlice() (error, []matrixStruct, int) {
	var count int
	var mtr []matrixStruct
	fmt.Print("Введите количество заявок:")
	_, err := fmt.Scanf("%d", &count)
	if err == nil {
		var size int
		fmt.Print("Введите размерность матрицы:")
		_, err = fmt.Scanf("%d", &size)
		if err == nil {
			err, mtr = createMatrixEx(count, size, 1)
		}
	}

	return err, mtr, count
}

func printMatrix(matrix [][]float64, size int) {
	for i := 0; i < size; i++ {
		fmt.Println(matrix[i])
	}
}

func printResult(slice []matrixStruct, count int) {
	for _, item := range slice {
		fmt.Println("Обратная матрица:")
		printMatrix(item.inverseMatrix, item.size)
		fmt.Println("Стандартный алгоритм умножения:")
		printMatrix(item.standartMul, item.size)
		fmt.Println("Алгоритм Винограа:")
		printMatrix(item.vinogradMul, item.size)
		fmt.Println()
	}
}

func main() {
	num := -1
	menu := "1)Запустить конвейер\n2)Последовательная обработка\n3)Замер времени\n0)Выход"
	for num != 0 {
		fmt.Println(menu)
		_, err := fmt.Scanf("%d", &num)
		if err != nil {
			num = -1
		}

		if num == 1 {
			err, matrixSlice, count := createSlice()
			if err != nil {
				continue
			}
			pipelineMtr(matrixSlice, count, 1)
			printResult(matrixSlice, count)
		} else if num == 2 {
			err, matrixSlice, count := createSlice()
			if err != nil {
				continue
			}
			serialMtr(matrixSlice, count, 1)
		} else if num == 3 {
			checkTime()
		} else if num == 0 {
			return
		} else {
			fmt.Println("Данные введены неверно")
		}
	}
}
