package pkg

import (
	"fmt"
	"math"
	_ "strconv"
	"strings"
)

func ShowTable(rows [][]float64) {
	colWidths := []int{}
	rowsRounded := [][]float64{}
	for i, _ := range rows {
		row := &rows[i]
		outputRow := []float64{}
		for j, _ := range *row {
			rounded := roundFloat((*row)[j], 3)
			var vStr string
			if math.Round(rounded) == rounded {
				vStr = fmt.Sprintf("%.0f", rounded)
			} else {
				vStr = fmt.Sprintf("%.3f", rounded)
			}
			if len(colWidths) < j+1 {
				// fmt.Println("Using " + strconv.Itoa(len(vStr)) + " as size of j=" + strconv.Itoa(j) + " because vStr=" + vStr)
				colWidths = append(colWidths, len(vStr))
			} else if colWidths[j] < len(vStr) {
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
