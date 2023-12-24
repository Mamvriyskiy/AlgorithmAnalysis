package main

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
	"github.com/olekukonko/tablewriter"
)


func checkTime() error {
	var count int
	var size int
	_, err := fmt.Scanf("%d", &count)
	if err == nil {
		_, err = fmt.Scanf("%d", &size)
	}

	if err != nil {
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Размерность",  "Конвейер", "Последовательный"})

	num := 0
	err, matrixSlice := createMatrixEx(1, size, 0)
	for i := 1; i <= 1000; i+= num {
		startA := C.getThreadCpuTimeNs()
		for j := 0; j < 5; j++ {
			pipelineMtr(matrixSlice, i, 0)
		}
		finishA := C.getThreadCpuTimeNs()
		timeA := int64(finishA - startA) / 5

		startB := C.getThreadCpuTimeNs()
		for j := 0; j < 5; j++ {
			serialMtr(matrixSlice, i, 0)
		}
		finishB := C.getThreadCpuTimeNs()
		timeB := int64(finishB - startB) / 5
		
		table.Append([]string{fmt.Sprintf("%d", i), fmt.Sprintf("%d", timeA), fmt.Sprintf("%d", timeB)})

		fmt.Println("Completed:", i)

		if i == 1 {
			num = 9
		} else if i == 10 {
			num = 40
		} else if i == 50 {
			num = 50
		} else if i == 100 {
			num = 100
		}
	}

	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.SetColumnSeparator("|")
	table.SetRowSeparator("-")
	table.SetColWidth(20)
	table.SetAutoWrapText(false)

	table.Render()
    fmt.Println()

	return err
}
