package pkg

import (
	"math"
	"strconv"
)

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func sciNotation(input float64) string {
	count := 0
	for input >= 10 {
		count += 1
		input = input / 10.0
	}
	return strconv.FormatFloat(input, 'e', -1, 64) + "e" + strconv.Itoa(count)
}

func factorial(input float64) (ret float64) {
	ret = 1
	for n := float64(1); n <= input; n++ {
		ret = ret * n
	}
	return ret
}

func factorialMin(input float64) float64 {
	return math.Pow(2*math.Pi*input, 0.5) * math.Pow(input/math.E, input) *
		math.Pow(math.E, 1/((12*input)+1))
}

func factorialMax(input float64) float64 {
	return math.Pow(2*math.Pi*input, 0.5) * math.Pow(input/math.E, input) *
		math.Pow(math.E, 1/(12*input))
}
