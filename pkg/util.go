package pkg

import (
	"fmt"
	"math"
	_ "strconv"
	"strings"
)

type Table struct {
	Headers []string
	Rows [][]float64
}

func ShowTable(table *Table) {
	colWidths := []int{}	
	for i, _ := range table.Headers {
		colWidths = append(colWidths, len(table.Headers[i]))
	}
	
	rowsRounded := [][]float64{}
	for i, _ := range table.Rows {
		row := &table.Rows[i]
		outputRow := []float64{}
		for j, _ := range *row {
			rounded := roundFloat((*row)[j], 3)
			var vStr string
			if math.Round(rounded) == rounded {
				vStr = fmt.Sprintf("%.0f", rounded)
			} else {
				vStr = fmt.Sprintf("%.3f", rounded)
			}
			if colWidths[j] < len(vStr) {
				// fmt.Println("Using " + strconv.Itoa(len(vStr)) + " as size of j=" + strconv.Itoa(j))
				colWidths[j] = len(vStr)
			}
			outputRow = append(outputRow, rounded)
		}
		rowsRounded = append(rowsRounded, outputRow)
	}

	if len(rowsRounded) == 0 {
		return
	}

	fillers := []string{}
	for j, _ := range rowsRounded[0] {
		filler := strings.Repeat("_", colWidths[j])
		fillers = append(fillers, filler)
	}
	
	headerLine := "|"
	for i, _ := range table.Headers {
		headerLine += strings.Repeat(" ", len(fillers[i]) - len(table.Headers[i])) + table.Headers[i] + "|"
	}
	fmt.Println(headerLine)

	fmt.Println("|" + strings.Join(fillers, "|") + "|")
	for i, _ := range rowsRounded {
		row := &rowsRounded[i]
		fmt.Print("|")
		for j, _ := range *row {
			v := (*row)[j]
			if math.Round(v) == v {
				fmt.Printf("%*.0f", colWidths[j], v)
			} else {
				fmt.Printf("%.3f", v)
			}
			fmt.Print("|")
		}
		fmt.Println()
	}
}
