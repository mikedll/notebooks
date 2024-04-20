package pkg

import (
	"fmt"
	"math"
)

func Exercise3p3d5b() *TwoLines {
	expr1 := func(n float64) float64 {
		return roundFloat(math.Pow(n, 3), 3)
	}

	expr2 := func(n float64) float64 {
		factArg := math.Ceil(math.Log2(math.Log2(n)))
		return roundFloat(factorialMin(factArg), 3)
	}

	result := &TwoLines{}
	for n := float64(1); n <= 30; n++ {
		v1 := expr1(n)
		v2 := expr2(n)
		result.X = append(result.X, n)
		result.Y = append(result.Y, v1)
		result.Y2 = append(result.Y2, v2)
	}

	return result
}

func Exercise3p3d5a() *TwoLines {
	expr1 := func(n float64) float64 {
		factArg := math.Ceil(math.Log2(n))
		return roundFloat(float64(factorialMin(factArg)), 3)
	}

	expr2 := func(n float64) float64 {
		return roundFloat(math.Pow(n, 5), 3)
	}

	result := &TwoLines{}

	start := math.Pow(float64(2), 78)
	end := math.Pow(float64(2), 88)
	step := (end - start) / 10

	for n := start; n <= end; n += step {
		v1 := expr1(n)
		v2 := expr2(n)
		result.X = append(result.X, n)
		result.Y = append(result.Y, v1)
		result.Y2 = append(result.Y2, v2)
	}

	examined := []float64{
		math.Pow(float64(2), 78),
		math.Pow(float64(2), 80),
		math.Pow(float64(2), 82),
		math.Pow(float64(2), 84),
		math.Pow(float64(2), 86),
		math.Pow(float64(2), 88),
	}
	for _, n := range examined {
		v1 := expr1(n)
		v2 := expr2(n)

		s := fmt.Sprintf("|%s|%s|%s|", sciNotation(n), sciNotation(v1), sciNotation(v2))
		// s := fmt.Sprintf("|%20.2f|%20.f|%20.f|", n, v1, v2)
		fmt.Println(s)
	}

	return result
}

func Exercise3p3d4c() *TwoLines {
	expr1 := func(c5, c6 float64) float64 {
		return c5 / c6
	}
	expr2 := func(ratio float64) float64 {
		return math.Pow(ratio, 0.25)
	}
	expr3 := func(n, ratio float64) float64 {
		return ratio * n
	}
	expr4 := func(originalInput float64) float64 {
		return math.Log2(originalInput)
	}
	expr5 := func(n float64) float64 {
		return 5 * math.Log2(n)
	}

	inputs := []float64{10, 12, 14, 16, 18, 20}
	rows := [][]float64{}
	for _, n := range inputs {
		ratio := expr1(50_000, 1)
		cutoff := expr2(ratio)
		originalInput := expr3(n, ratio)
		v1 := expr5(n)
		v2 := expr4(originalInput)
		row := []float64{ratio, cutoff, originalInput, v1, v2}
		rows = append(rows, row)
	}
	ShowTable(rows)

	result := &TwoLines{}

	for n := float64(1); n <= 40; n += 1 {
		v1 := expr5(n)
		v2 := expr4(n)
		result.X = append(result.X, n)
		result.Y = append(result.Y, roundFloat(v1, 2))
		result.Y2 = append(result.Y2, roundFloat(v2, 2))
	}

	return result
}
