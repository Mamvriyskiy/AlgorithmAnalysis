package main

import (
	"fmt"
	"os"
	"bufio"
	"./algorithms"
	"./checktime"
	"strconv"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("No argument")
		os.Exit(-1)
	}

	num := args[0]

	numV, err := strconv.Atoi(num)

	if err != nil {
		os.Exit(-1)
	}

	if numV == 1 {
		mulEx()
	} else {
		checktime.CheckTime()
	}

}

func mulEx() {
	fmt.Print("Size A: ")
	matrixA, aA, bA := inputMatrixEX()

	fmt.Print("Size B: ")
	matrixB, aB, bB := inputMatrixEX()

	if bA != aB {
		os.Exit(0)
	}

	fmt.Println("MatrixA:")
	printMatrix(matrixA, aA, bA)
	fmt.Println()
	fmt.Println("MatrixB:")
	printMatrix(matrixB, aB, bB)

	//Простой способ умножения 
	fmt.Println()
	fmt.Println("Standart:")
	matrixC := algorithms.MulMAtrixA(matrixA, matrixB, aA, bB, bA)
	printMatrix(matrixC, aA, bB)

	//Оптимизированный простой способ умножения
	fmt.Println()
	fmt.Println("Winograd:")
	matrixD := algorithms.MulMAtrixB(matrixA, matrixB, aA, bB, bA)
	printMatrix(matrixD, aA, bB)

	//алгоритм Виноградова
	fmt.Println()
	fmt.Println("WinogradOpt:")
	matrixF := algorithms.MulVinogradovA(matrixA, matrixB, aA, bA, bB)
	printMatrix(matrixF, aA, bB)
}


func printMatrix(matrix [][]int, a, b int) {
	for i := 0; i < a; i++ {
		for j := 0; j < b; j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println()
	}
}

func inputMatrixEX() ([][]int, int, int) {
	reader := bufio.NewReader(os.Stdin)

	a, b := inputSizeMatrix(reader)
	fmt.Println()

	matrix := createMatrix(reader, a, b)

	return matrix, a, b
}

func createMatrix(reader *bufio.Reader, a, b int) [][]int {
	matrix := make([][]int, a)

	for i := 0; i < a; i++ {
		matrix[i] = make([]int, b)
		for j := 0; j < b; j++ {
			fmt.Printf("[%d][%d]: ", i + 1, j + 1)
			_, err := fmt.Fscan(reader, &matrix[i][j])
			if err != nil {
				os.Exit(-1)
			}
		}

		fmt.Println()
	}

	return matrix
}


func inputSizeMatrix(reader *bufio.Reader) (int, int) {
	var a, b int
	_, err := fmt.Fscan(reader, &a, &b)
	if err != nil {
		os.Exit(-1)
	}

	return a, b
}

