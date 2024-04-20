package pkg

import (
	"fmt"
	"math"
)

func Exercise3p3d4c() *ThreeLines {
	expr1 := func(n float64) float64 {
		return 5 * math.Log2(n)
	}
	expr2 := func(n float64) float64 {
		return math.Log2(50000.0 * n)
	}
	expr3 := func(n float64) float64 {
		inputN := n + 10
		if inputN <= 0 {
			return 0
		}
		return math.Log2(50000.0 * inputN)
	}

	result := &ThreeLines{}
	for n := float64(1); n <= 40; n += 1 {
		v1 := expr1(n)
		v2 := expr2(n)
		v3 := expr3(n)
		result.X = append(result.X, n)
		result.Y = append(result.Y, roundFloat(v1, 2))
		result.Y2 = append(result.Y2, roundFloat(v2, 2))
		result.Y3 = append(result.Y3, roundFloat(v3, 2))
	}

	return result
}

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
