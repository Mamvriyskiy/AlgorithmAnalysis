package main

import (
	"bufio"
	"fmt"
	"os"

	"./algorithms"
	"./checktime"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	num := 0
	var err error
	for {
		fmt.Println("Меню:\n1)Поиск дистанции;\n2)Проверка времени;\n3)Выход.\nВвод: ")
		_, err = fmt.Fscan(reader, &num)
		if err != nil {
			os.Exit(-1)
		}

		fmt.Println()
		if num == 1 {
			searchDistance(reader)
		} else if num == 2 {
			checktime.CheckTime()
			checktime.CheckMemoryEx()
		} else {
			break
		}	
		fmt.Println()
	}
}

func searchDistance(reader *bufio.Reader) {
	var s1, s2 string
	_, err := fmt.Fscan(reader, &s1)
	if err != nil {
		os.Exit(-1)
	}

	_, err = fmt.Fscan(reader, &s2)
	if err != nil {
		os.Exit(-1)
	}

	r1 := []rune(s1)
	r2 := []rune(s2)

	//матричный алгоритм Левенштейна
	fmt.Println("Матрица Левенштейна:", algorithms.MatrixLevenshtein(r1, r2))

	//матричный алгоритм Дамеруа-Левенштейна
	fmt.Println("Матрица Дамерау-Левенштейна:", algorithms.MatrixDamerauLevenshtein(r1, r2))

	//рекурсивный алгоритм Дамеруа-Левенштейна
	fmt.Println("Рекурсивный алгоритм Дамерау-Левенштейна:", algorithms.RecursiveDL(r1, r2))

	//рекурсивный алгоритм Дамеруа-Левенштейна с кешем
	fmt.Println("Рекурсивный алгоритм Дамерау-Левенштейна с кешем:", algorithms.RecursiveDLCash(r1, r2))
}