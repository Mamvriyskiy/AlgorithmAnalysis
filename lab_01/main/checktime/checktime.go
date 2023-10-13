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
	"math/rand"
	"../algorithms"
	"github.com/olekukonko/tablewriter"
)

const MAXLENGHTSTR = 15

func CheckTime() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Lenght", "MatrxL", "MatrixDL", "RCDL", "RDL"})

	STEPLENGHT := 1
    for i := STEPLENGHT; i <= MAXLENGHTSTR; i += STEPLENGHT {
        s1 := generateString(i)
        s2 := generateString(i)

        var timeMatrixL int64
        var timeMatrixDL int64
        var timeRCashDL int64
        var timeRDL int64

        start := C.getThreadCpuTimeNs()
        for j := 0; j < 10; j++ {
            algorithms.MatrixLevenshtein(s1, s2)
        }
        finish := C.getThreadCpuTimeNs()
        timeMatrixL = int64(finish - start) / 10

        start = C.getThreadCpuTimeNs()
        for j := 0; j < 10; j++ {
            algorithms.MatrixDamerauLevenshtein(s1, s2)
		}
        finish = C.getThreadCpuTimeNs()
        timeMatrixDL = int64(finish - start) / 10

        start = C.getThreadCpuTimeNs()
        for j := 0; j < 10; j++ {
            algorithms.RecursiveDLCash(s1, s2)
        }
        finish = C.getThreadCpuTimeNs()
        timeRCashDL = int64(finish - start)

        if i <= 15 {
            start := C.getThreadCpuTimeNs()
            algorithms.RecursiveDL(s1, s2)
            finish := C.getThreadCpuTimeNs()
            timeRDL = int64(finish - start)
        }

		if i <= 15 {
        	table.Append([]string{fmt.Sprintf("%d", i), fmt.Sprintf("%d", timeMatrixL), fmt.Sprintf("%d", timeMatrixDL), fmt.Sprintf("%d", timeRCashDL), fmt.Sprintf("%d", timeRDL)})
		} else {
			table.Append([]string{fmt.Sprintf("%d", i), fmt.Sprintf("%d", timeMatrixL), fmt.Sprintf("%d", timeMatrixDL), fmt.Sprintf("%d", timeRCashDL), fmt.Sprintf("-")})
		}

        // Settings style table
        table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
        table.SetCenterSeparator("|")
        table.SetColumnSeparator("|")
        table.SetRowSeparator("-")
		table.SetColWidth(20)
        table.SetAutoWrapText(false)

		if i == 15 {
			STEPLENGHT = 30
		}

        if i == 100 {
            STEPLENGHT = 100
        }
            
    }

	table.Render()
}


func generateString(length int) []rune {
    letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
    b := make([]rune, length)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }

	return b
}