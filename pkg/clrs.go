package pkg

import (
	"fmt"
	"math"
)

func Exercise3p3d5b() *TwoLines {
	expr1 := func(n float64) float64 {
		return roundFloat(math.Pow(n, 2), 3)
	}

	expr2 := func(n float64) float64 {
		factArg := math.Ceil(math.Log2(math.Log2(n)))
		return roundFloat(factorialMax(factArg), 3)
	}

	// start := float64(1)
	start := math.Pow(float64(2), 88)
	end := math.Pow(float64(2), 150)
	step := (end - start) / 200

	result := &TwoLines{}
	for n := start; n <= end; n += step {
		v1 := expr2(n)
		v2 := expr1(n)
		// fmt.Println("Iterating ", start, n, end, v1, v2)

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
		// return roundFloat(math.Pow(n, 5), 3)
		return roundFloat(math.Pow(n, 5), 3)
	}

	result := &TwoLines{}

	start := math.Pow(float64(2), 0)
	end := math.Pow(float64(2), 88)
	// start := math.Pow(float64(2), 0)
	// end := math.Pow(float64(2), 14)
	// end := math.Pow(float64(2), 8)

	step := (end - start) / 200

	for n := start; n <= end; n += step {
		fmt.Println("Step: ", step)
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

func Exercise3p3d4c() *FourLines {
	expr1 := func(c6, c7 float64) float64 {
		return c6 / c7
	}
	expr2 := func(ratio float64) float64 {
		return math.Pow(ratio, 0.25)
	}
	expr3 := func(ratio, n float64) float64 {
		return ratio * n
	}
	expr4 := func(originalInput float64) float64 {
		return math.Log2(originalInput)
	}
	expr5 := func(n float64) float64 {
		return 5 * math.Log2(n)
	}
	expr6 := func(c3, c4 float64) float64 {
		return c3 / c4
	}
	expr7 := func(c3, c4 float64) float64 {
		return c4 / c3
	}
	expr8 := func(ratio, n float64) float64 {
		return ratio * n
	}
	expr9 := func(originalInput float64) float64 {
		return math.Log2(originalInput)
	}
	expr10 := func(n float64) float64 {
		return 0.5 * math.Log2(n)
	}
	expr11 := func(ratio float64) float64 {
		return math.Pow(ratio, 2)
	}

	c3 := float64(1.0)
	c4 := float64(10.0)
	c6 := float64(50_000)
	c7 := float64(1)
	fmt.Printf("Upper bound transition point: %.3f\n", expr2(expr1(c6, c7)))
	fmt.Printf("Lower bound transition point: %.3f\n", expr11(expr7(1, 10)))

	inputs := []float64{10, 12, 14, 16, 18, 20}
	rows := [][]float64{}
	for _, n := range inputs {
		ratio := expr1(50_000, 1)
		cutoff := expr2(ratio)
		originalInput := expr3(ratio, n)
		v1 := expr5(n)
		v2 := expr4(originalInput)
		row := []float64{ratio, cutoff, originalInput, v1, v2}
		rows = append(rows, row)
	}
	table := &Table{}
	table.Headers = []string{"ratio", "cutoff", "originalInput", "v1", "v2"}
	table.Rows = rows
	ShowTable(table)

	result := &FourLines{}

	lBoundRows := [][]float64{}
	for n := float64(1); n <= 400; n += 1 {
		v1 := expr10(n)
		v2 := expr9(expr8(expr6(c3, c4), n))

		row := []float64{expr6(c3, c4), expr8(expr6(c3, c4), n), expr9(expr8(expr6(c3, c4), n)), expr10(n)}
		lBoundRows = append(lBoundRows, row)

		v3 := expr5(n)
		v4 := expr4(expr3(expr1(c6, c7), n))
		result.X = append(result.X, n)
		result.Y = append(result.Y, roundFloat(v3, 2))
		result.Y2 = append(result.Y2, roundFloat(v4, 2))
		result.Y3 = append(result.Y3, roundFloat(v1, 2))
		result.Y4 = append(result.Y4, roundFloat(v2, 2))
	}
	// ShowTable(lBoundRows)

	return result
}

func Exercise3p3d6() {
	logStar := func(n float64) (int, float64) {
		i := 0
		for n > 1 {
			n = math.Log2(n)
			i += 1
		}
		return i, n
	}
	val, n := logStar(math.Pow(2, 3))
	
	
	fmt.Println("val, n", val, n)
	
	g_i := func(i int, n float64) float64 {
		for j := 1; j <= i; j++ {
			n = math.Pow(2, n)
		}
		return n
	}
	
	val2 := g_i(val, n)
	fmt.Println("val2", val2)	
	return
	
	table := &Table{}
	table.Headers = []string{ "n", "lg n" }
	inputs := []float64 { 2, 4, 16, math.Pow(2, 16) }
	for _, n := range inputs {
		result := []float64 { n, math.Log2(n) }
		table.Rows = append(table.Rows, result)
	}
	
	result := []float64 { 1, math.Pow(2, 16) }
	table.Rows = append(table.Rows, result)
	
	ShowTable(table)
	
}

func Exercise3p3d9() *TwoLines {
	result := &TwoLines{}
	
	expr1 := func(n float64) float64 {
		return n * math.Log2(n)
	}
	expr2 := func(n float64) float64 {
		return n;
	}
	
	
	for n := float64(1); n <= 50; n += 1 {
		result.X = append(result.X, n)
		result.Y = append(result.Y, expr1(n))
		result.Y2 = append(result.Y2, expr2(n))
	}
	
	return result
}

func Problem3d2() *TwoLines {
	result := &TwoLines{}
	
	expr1 := func(n float64) float64 {
		return math.Pow(n, 0.5);
	}
	
	expr2 := func(n float64) float64 {
		return math.Pow(n, math.Sin(n));
	}
	
	for n := float64(1); n<=200; n += 1 {
		result.X = append(result.X, n)
		result.Y = append(result.Y, expr1(n))
		result.Y2 = append(result.Y2, expr2(n))
	}
	
	return result
}

func Problem3d3() {
	itr := func(n float64) int {
		i := 0
		for n > 1 {
			n = math.Log2(n)
			i += 1
		}
		return i
	}
	
	inputs := []float64{ float64(4), float64(5), float64(8), 
		float64(16), float64(17), float64(2000), math.Pow(2, 16), 
		math.Pow(2, 40),  math.Pow(2,80), 
	}
	
	table := &Table{}
	table.Headers = []string{ "n", "lg^* n", "2^{lg^* n}", "lg n"}
	for _, input := range inputs {
		itr := float64(itr(input))
		twoPower := math.Pow(2, itr)
		row := []float64{ input, itr, twoPower, math.Log2(input) }
		table.Rows = append(table.Rows, row)
	}

  ShowTable(table)
}
