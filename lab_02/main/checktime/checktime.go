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
		// return t.tv_sec * 1000000000LL + t.tv_nsec; //нано
        //return t.tv_sec * 1000000LL + t.tv_nsec / 1000LL; //микро
        return t.tv_sec * 1000LL + t.tv_nsec / 1000000LL; // милл
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
	"../algorithms"
    //"runtime"
	"github.com/olekukonko/tablewriter"
)

const MAXLENGHTSTR = 1001

func CheckTime() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Длина", "Стандартный", "Виноград", "ВиноградОпт"})

    tableM := tablewriter.NewWriter(os.Stdout)
	tableM.SetHeader([]string{"Длина", "Стандартный", "Виноград", "ВиноградОпт"})

    lenght := []int{}
    a := []int{}
    c := []int{}
    d := []int{}

    START := 10
	STEPLENGHT := 1

    for i := START; i <= MAXLENGHTSTR; i += STEPLENGHT {
        m1 := generateMatrix(i)
        m2 := generateMatrix(i)

        var timeA int64
        var timeC int64
        var timeD int64

        // var memoryA int64
        // var memoryC int64
        // var memoryD int64

        // {
        //     start := C.getThreadCpuTimeNs()
        //     algorithms.MulMAtrixA(m1, m2, i, i, i)
        //     finish := C.getThreadCpuTimeNs()
        //     timeA = int64(finish - start)
        //}


        {
            start := C.getThreadCpuTimeNs()
            for j := 0; j < 5; j++ {
                algorithms.MulMAtrixA(m1, m2, i, i, i)
            }
            finish := C.getThreadCpuTimeNs()
            timeA = int64(finish - start) / 5
        }
        
        // var b1, b2 runtime.MemStats
        // runtime.ReadMemStats(&b1)
        // algorithms.MulMAtrixA(m1, m2, i, i, i)
        // runtime.ReadMemStats(&b2)
        // memoryA = int64(b2.TotalAlloc - b1.TotalAlloc)

        // {
        //     start := C.getThreadCpuTimeNs()
        //     algorithms.MulMAtrixB(m1, m2, i, i, i)
        //     finish := C.getThreadCpuTimeNs()
        //     timeB = int64(finish - start)
		// }

        {
            start := C.getThreadCpuTimeNs()
            for j := 0; j < 5; j++ {
                algorithms.MulVinogradovA(m1, m2, i, i, i)
            }
            finish := C.getThreadCpuTimeNs()
            timeC = int64(finish - start) / 5
        }

        // runtime.ReadMemStats(&b1)
        // algorithms.MulVinogradovA(m1, m2, i, i, i)
        // runtime.ReadMemStats(&b2)
        // memoryC = int64(b2.TotalAlloc - b1.TotalAlloc)

        {
            start := C.getThreadCpuTimeNs()
            for j := 0; j < 5; j++ {
                algorithms.MulVinogradovB(m1, m2, i, i, i)
            }
            finish := C.getThreadCpuTimeNs()
            timeD = int64(finish - start) / 5
        }

        // runtime.ReadMemStats(&b1)
        // algorithms.MulVinogradovB(m1, m2, i, i, i)
        // runtime.ReadMemStats(&b2)
        // memoryD = int64(b2.TotalAlloc - b1.TotalAlloc)

        // {
        //     start := C.getThreadCpuTimeNs()
        //     algorithms.MulVinogradovB(m1, m2, i, i, i)
        //     finish := C.getThreadCpuTimeNs()
        //     timeD = int64(finish - start)
		// }

        lenght = append(lenght, i)
        a = append(a, int(timeA))
        d = append(d, int(timeD))
        c = append(c, int(timeC))
		table.Append([]string{fmt.Sprintf("%d", i), fmt.Sprintf("%d", timeA), fmt.Sprintf("%d", timeC), fmt.Sprintf("%d", timeD)})
        // tableM.Append([]string{fmt.Sprintf("%d", i), fmt.Sprintf("%d", memoryA), fmt.Sprintf("%d", memoryC), fmt.Sprintf("%d", memoryD)})

        // Settings style table
        table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
        table.SetCenterSeparator("|")
        table.SetColumnSeparator("|")
        table.SetRowSeparator("-")
		table.SetColWidth(20)
        table.SetAutoWrapText(false)
        
        // tableM.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
        // tableM.SetCenterSeparator("|")
        // tableM.SetColumnSeparator("|")
        // tableM.SetRowSeparator("-")
		// tableM.SetColWidth(20)
        // tableM.SetAutoWrapText(false)

        if i == 10 {
            STEPLENGHT = 90
        }

        if i == 100 {
            STEPLENGHT = 100
        }

        fmt.Println("Complite: ", i)
    }

    createPythonScript(lenght, a, c, d)
	table.Render()
    fmt.Println()
    // tableM.Render()
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

func generateMatrix(length int) [][]int {
    matrix := make([][]int, length)

    for i := 0; i < length; i++ {
        matrix[i] = make([]int, length)
        for j := 0; j < length; j++ {
            matrix[i][j] = rand.Intn(1000)
        }
    }

	return matrix
}
