package checktime

/*
	#include <pthread.h>
	#include <time.h>
	#include <stdio.h>

	static long long getThreadCpuTimeNs() {
		struct timespec t;
		if (clock_gettime(CLOCK_MONOTONIC, &t)) {
			perror("clock_gettime");
			return 0;
		}
		return t.tv_sec * 1000000000LL + t.tv_nsec;
	}
*/
import "C"
import (
	"fmt"
	"os"
    "log"
    "strings"
    "os/exec"
	"math/rand"
	"../sorted"
	"github.com/olekukonko/tablewriter"
)

type createSliceFunc func(int) []int

const MAXLENGHTSTR = 5000

func CheckTimeEX() {
	fmt.Println("Отсортированный по возрастанию:")
	checkTime(generateSliceAsc)
	fmt.Println("Отсортированный по убыванию:")
	checkTime(generateSliceDesc)
	fmt.Println("Рандомный слайс:")
	checkTime(generateSliceRand)
}

func checkTime(createFunc createSliceFunc) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Lenght", "Pancake", "HeapSort", "ShakerSort"})

    lenght := []int{}
    a := []int{}
    c := []int{}
    d := []int{}

    START := 10
	STEPLENGHT := 1

    for i := START; i <= MAXLENGHTSTR; i += STEPLENGHT {
        slice := createFunc(i)

        var timeA int64
        var timeC int64
        var timeD int64

        {
            start := C.getThreadCpuTimeNs()
            for j := 0; j < 10; j++ {
                sorted.PancakeSort(slice)
            }
            finish := C.getThreadCpuTimeNs()
            timeA = int64(finish - start) / 10
        }
    
        {
            start := C.getThreadCpuTimeNs()
            for j := 0; j < 10; j++ {
                sorted.HeapSort(slice)
            }
            finish := C.getThreadCpuTimeNs()
            timeC = int64(finish - start) / 10
        }

        {
            start := C.getThreadCpuTimeNs()
            for j := 0; j < 10; j++ {
                sorted.ShakerSort(slice)
            }
            finish := C.getThreadCpuTimeNs()
            timeD = int64(finish - start) / 10
        }


        lenght = append(lenght, i)
        a = append(a, int(timeA))
        d = append(d, int(timeD))
        c = append(c, int(timeC))
		table.Append([]string{fmt.Sprintf("%d", i), fmt.Sprintf("%d", timeA), fmt.Sprintf("%d", timeC), fmt.Sprintf("%d", timeD)})

        // Settings style table
        table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
        table.SetCenterSeparator("|")
        table.SetColumnSeparator("|")
        table.SetRowSeparator("-")
		table.SetColWidth(20)
        table.SetAutoWrapText(false)
        
        if i == 10 {
            STEPLENGHT = 90
        }

        if i == 100 {
            STEPLENGHT = 100
        }

		if i == 1000 {
            STEPLENGHT = 1000
        }

        fmt.Println("Complite: ", i)
    }

    createPythonScript(lenght, a, c, d)
	table.Render()
}
   
func createPythonScript(length, a, c, d []int) {
    lengthStrings := make([]string, len(length))
    for i, num := range length {
        lengthStrings[i] = fmt.Sprintf("%d", num)
    }

    aStrings := make([]string, len(a))
    for i, num := range a {
        aStrings[i] = fmt.Sprintf("%d", num)
    }

    cStrings := make([]string, len(c))
    for i, num := range c {
        cStrings[i] = fmt.Sprintf("%d", num)
    }

    dStrings := make([]string, len(d) + 1)
    for i, num := range d {
        dStrings[i] = fmt.Sprintf("%d", num)
    }

    dStrings[len(d)] = fmt.Sprintf("%d", len(d))
     
    // Команда Python, которую вы хотите выполнить, с передачей аргументов
    pythonScript := "./checktime/shema.py"
    pythonArgs := append([]string{pythonScript}, lengthStrings...)
    pythonArgs = append(pythonArgs, aStrings...)
    pythonArgs = append(pythonArgs, cStrings...)
    pythonArgs = append(pythonArgs, dStrings...)

    // Создание команды
    cmd := exec.Command("python3", pythonArgs...)

    var stdout, stderr strings.Builder
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr

    // Запуск команды и ожидание ее завершения
    err := cmd.Run()
    if err != nil {
       log.Fatalf("Ошибка выполнения Python-скрипта: %v\n", err)
    }

    result := stdout.String()
    errorOutput := stderr.String()
    fmt.Printf("Стандартный вывод Python-скрипта: %s\n", result)
    fmt.Printf("Стандартная ошибка Python-скрипта: %s\n", errorOutput)
}

func generateSliceAsc(length int) []int {
    slice := make([]int, length)

    for i := 0; i < length; i++ {
        slice[i] = i

    }

	return slice
}

func generateSliceDesc(length int) []int {
    slice := make([]int, length)

	n := 0
    for i := length - 1; i > 0; i-- {
        slice[n] = i
		n++
    }

	return slice
}

func generateSliceRand(length int) []int {
    slice := make([]int, length)

    for i := 0; i > 0; i++ {
        slice[i] = rand.Intn(1000)

    }

	return slice
}

