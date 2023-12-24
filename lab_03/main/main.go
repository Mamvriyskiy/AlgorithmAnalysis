package main

import (
	"fmt"
	"bufio"
	"os"
	"./sorted"
	"./checktime"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var err error

	for err == nil {
		var n int
		fmt.Println("Выберите пункт меню:\n1)Одиночный случай;\n2)Запуск тестов;\n3)Выход из программы.")
		_, err = fmt.Fscan(reader, &n)

		if n == 1 {
			fmt.Println()
			lenl, slice := createSlice(reader)
			fmt.Println()

			slice1 := make([]int, lenl)
			slice2 := make([]int, lenl)
			slice3 := make([]int, lenl)

			//
			copy(slice1, slice)
			sorted.HeapSort(slice1)
			fmt.Print("Пирамидальная сортировка : ")
			outputSlice(slice1)
			fmt.Println()
			//
			copy(slice2, slice)
			sorted.PancakeSort(slice2)
			fmt.Print("Блинная сортирвока: ")
			outputSlice(slice2)
			fmt.Println()
			//
			copy(slice3, slice)
			sorted.ShakerSort(slice3)
			fmt.Print("Сортировка перемешиванием: ")
			outputSlice(slice3)
			fmt.Println()
		} else if n == 2 {
			checktime.CheckTimeEX()
		} else {
			os.Exit(0)
		}
		fmt.Println()
	}
}

func outputSlice(slice []int) {
	n := len(slice)
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", slice[i])
	}
}

func createSlice(reader *bufio.Reader) (int, []int) {
	var err error
	var n int
	fmt.Print("Введите размер слайса: ")
	_, err = fmt.Fscan(reader, &n)
	fmt.Println()

	slice := make([]int, n)

	for i := 0; i < n && err == nil; i++ {
		fmt.Printf("Введите %d элемент слайса: ", i + 1)
		_, err = fmt.Fscan(reader, &slice[i])
		if err != nil {
			os.Exit(-1)
		}
	}

	return n, slice
}
