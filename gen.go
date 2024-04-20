package main

import (
  "math"
  "log"
  "encoding/json"
  "io/ioutil"
  "github.com/dougwatson/Go/v3/math/factorial"
)

type TwoLines struct {
  X []int       `json:"x"` 
  Y []float64    `json:"y"`
  Y2 []float64   `json:"y2"`
}

type ThreeLines struct {
  X []int       `json:"x"` 
  Y []float64    `json:"y"`
  Y2 []float64   `json:"y2"`
  Y3 []float64   `json:"y3"`
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func exercise3p3d4c() *ThreeLines {
  expr1 := func(n int) float64 {
    return 5 * math.Log2(float64(n))
  }
  expr2 := func(n int) float64 {
    return math.Log2(50000.0 * float64(n))
  }
  expr3 := func(n int) float64 {
    inputN := n + 10;
    if(inputN <= 0) {
      return 0;
    }
    return math.Log2(50000.0 * float64(inputN))
  }
  
  result := &ThreeLines{}
  for n := 1; n <= 40; n += 1 {
    v1 := expr1(n)
    v2 := expr2(n)
    v3 := expr3(n)
    result.X = append(result.X, n)
    result.Y = append(result.Y, roundFloat(v1, 2))
    result.Y2 = append(result.Y2, roundFloat(v2, 2))
    result.Y3 = append(result.Y3, roundFloat(v3, 2))
  }
  
  return result;
}

func exercise3p3d5b() *TwoLines {
  expr1 := func(n int) float64 {
    return roundFloat(math.Pow(float64(n), 3), 3);
  }
  
  expr2 := func(n int) float64 {
    factArg := int(math.Ceil(math.Log2(math.Log2(float64(n)))))
    return roundFloat(float64(factorial.Iterative(factArg)), 3)
  }
  
  result := &TwoLines{}
  for n := 1; n <= 30; n++ {
    v1 := expr1(n);
    v2 := expr2(n);
    result.X = append(result.X, n)
    result.Y = append(result.Y, v1)
    result.Y2 = append(result.Y2, v2)
  }
  
  return result;
}

func main() {
  result := exercise3p3d5b();
  
  output, err := json.Marshal(*result)
  if err != nil {
    log.Fatal(err)
  }

  var outputPretty []byte
  outputPretty, err = json.MarshalIndent(*result, "", "  ")
  if err != nil {
    log.Fatal(err)
  }
  
  ioutil.WriteFile("data.json", output, 0644)
  log.Println("Wrote data.json")
  println(string(outputPretty))
}
