package checktime

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
	"unsafe"
)

func getSizeLevMtr(len int) int {
	return (len + 1) * (len + 1) * int(unsafe.Sizeof(int(0))) +
		2*int(unsafe.Sizeof(string(""))) +
		2*int(unsafe.Sizeof(int(0))) +
		2*int(unsafe.Sizeof(int(0))) +
		int(unsafe.Sizeof((*int)(nil))) + (len + 1) * int(unsafe.Sizeof((*int)(nil)))
}

func getSizeDlMtr(len int) int {
	return (len + 1) * (len + 1) * int(unsafe.Sizeof(int(0))) +
		2*int(unsafe.Sizeof(string(""))) +
		2*int(unsafe.Sizeof(int(0))) +
		3*int(unsafe.Sizeof(int(0))) +
		int(unsafe.Sizeof((*int)(nil))) + (len + 1) * int(unsafe.Sizeof((*int)(nil)))
}

func getSizeDlRec(len int) int {
	return (len + len) * (2*int(unsafe.Sizeof(int(0))) + 3*int(unsafe.Sizeof(int(0))) + 2*int(unsafe.Sizeof(string(""))))
}

func getSizeDlRecCash(len int) int {
	return getSizeDlRec(len) +
		(len + 1) * (len + 1) * int(unsafe.Sizeof(int(0))) +
		int(unsafe.Sizeof((*int)(nil))) + (len + 1) * int(unsafe.Sizeof((*int)(nil)))
}

const MAXLENGHTSTR = 1000

func CheckMemoryEx() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Длина строк", "НерекурсивныйЛ", "НерекурсивныйДЛ", "Рекурсивный с кэшем", "Рекурсивный"})

	STEPLENGHT := 1
    for i := STEPLENGHT; i <= MAXLENGHTSTR; i += STEPLENGHT {
		var RDLC int
		var L int
		var DL int
		var RDL int
		
		RDL = getSizeDlRec(i)
		L = getSizeLevMtr(i)
		DL = getSizeDlMtr(i)
		RDLC = getSizeDlRecCash(i)

		table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
        table.SetCenterSeparator("|")
        table.SetColumnSeparator("|")
        table.SetRowSeparator("-")
		table.SetColWidth(20)
        table.SetAutoWrapText(false)

		table.Append([]string{fmt.Sprintf("%d", i), fmt.Sprintf("%d", L), fmt.Sprintf("%d", DL), fmt.Sprintf("%d", RDLC), fmt.Sprintf("%d", RDL)})

		if i == 10 {
			STEPLENGHT = 90
		}

        if i == 100 {
            STEPLENGHT = 100
        }
	}

	table.Render()
}
