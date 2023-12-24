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
		//return t.tv_sec * 1000000000LL + t.tv_nsec; //нано
        return t.tv_sec * 1000000LL + t.tv_nsec / 1000LL; //микро
	}
*/
import "C"
import (
	"fmt"
	"os"
	"math/rand"
    "runtime"
	"../algorithms"
	"github.com/olekukonko/tablewriter"
)

func CheckTime() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Длина", "НерекурсивныйЛ", "НерекурсивныйДЛ", "Рекурсивный с кэшем", "Рекурсивный"})

    tableM := tablewriter.NewWriter(os.Stdout)
	tableM.SetHeader([]string{"Длина", "НерекурсивныйЛ", "НерекурсивныйДЛ", "Рекурсивный с кэшем", "Рекурсивный"})

	STEPLENGHT := 1
    for i := STEPLENGHT; i <= MAXLENGHTSTR; i += STEPLENGHT {
        s1 := generateString(i)
        s2 := generateString(i)

        var timeMatrixL int64
        var timeMatrixDL int64
        var timeRCashDL int64
        var timeRDL int64

        var memoryMatrixL int64
        var memoryMatrixDL int64
        var memoryRCashDL int64
        var memoryRDL int64

        //Test 1
        start := C.getThreadCpuTimeNs()
        for j := 0; j < 10; j++ {
            algorithms.MatrixLevenshtein(s1, s2)
        }
        finish := C.getThreadCpuTimeNs()
        timeMatrixL = int64(finish - start) / 10

        var m1, m2 runtime.MemStats
        runtime.ReadMemStats(&m1)
        algorithms.MatrixLevenshtein(s1, s2)
        runtime.ReadMemStats(&m2)
        memoryMatrixL = int64(m2.TotalAlloc - m1.TotalAlloc)

        //Test 2
        start = C.getThreadCpuTimeNs()
        for j := 0; j < 10; j++ {
            algorithms.MatrixDamerauLevenshtein(s1, s2)
		}
        finish = C.getThreadCpuTimeNs()
        timeMatrixDL = int64(finish - start) / 10

        runtime.ReadMemStats(&m1)
        algorithms.MatrixDamerauLevenshtein(s1, s2)
        runtime.ReadMemStats(&m2)
        memoryMatrixDL = int64(m2.TotalAlloc - m1.TotalAlloc)
        
        //Test 3
        start = C.getThreadCpuTimeNs()
        for j := 0; j < 10; j++ {
            algorithms.RecursiveDLCash(s1, s2)
        }
        finish = C.getThreadCpuTimeNs()
        timeRCashDL = int64(finish - start) / 10

        runtime.ReadMemStats(&m1)
        algorithms.RecursiveDLCash(s1, s2)
        runtime.ReadMemStats(&m2)
        memoryRCashDL = int64(m2.TotalAlloc - m1.TotalAlloc)

        //Test 4
        if i <= 10 {
            start := C.getThreadCpuTimeNs()
            algorithms.RecursiveDL(s1, s2)
            finish := C.getThreadCpuTimeNs()
            timeRDL = int64(finish - start)

            runtime.ReadMemStats(&m1)
            algorithms.RecursiveDL(s1, s2)
            runtime.ReadMemStats(&m2)
            memoryRDL = int64(m2.TotalAlloc - m1.TotalAlloc)
        }

		if i <= 10 {
        	table.Append([]string{fmt.Sprintf("%d", i), fmt.Sprintf("%d", timeMatrixL), fmt.Sprintf("%d", timeMatrixDL), fmt.Sprintf("%d", timeRCashDL), fmt.Sprintf("%d", timeRDL)})
		} else {
			table.Append([]string{fmt.Sprintf("%d", i), fmt.Sprintf("%d", timeMatrixL), fmt.Sprintf("%d", timeMatrixDL), fmt.Sprintf("%d", timeRCashDL), fmt.Sprintf("-")})
		}

        if i <= 10 {
        	tableM.Append([]string{fmt.Sprintf("%d", i), fmt.Sprintf("%d", memoryMatrixL), fmt.Sprintf("%d", memoryMatrixDL), fmt.Sprintf("%d", memoryRCashDL), fmt.Sprintf("%d", memoryRDL)})
		} else {
			tableM.Append([]string{fmt.Sprintf("%d", i), fmt.Sprintf("%d", memoryMatrixL), fmt.Sprintf("%d", memoryMatrixDL), fmt.Sprintf("%d", memoryRCashDL), fmt.Sprintf("-")})
		}

        // Settings style table
        table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
        table.SetCenterSeparator("|")
        table.SetColumnSeparator("|")
        table.SetRowSeparator("-")
		table.SetColWidth(20)
        table.SetAutoWrapText(false)

        tableM.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
        tableM.SetCenterSeparator("|")
        tableM.SetColumnSeparator("|")
        tableM.SetRowSeparator("-")
		tableM.SetColWidth(20)
        tableM.SetAutoWrapText(false)

		if i == 10 {
			STEPLENGHT = 90
		}

        if i == 100 {
            STEPLENGHT = 100
        }
        fmt.Println("Complete:", i)   
    }

	table.Render()
    fmt.Println()
    tableM.Render()
}


func generateString(length int) []rune {
    letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
    b := make([]rune, length)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }

	return b
}