package main

import (
  "math"
  "log"
  "encoding/json"
  "io/ioutil"
)

type TwoLines struct {
  X []int       `json:"x"` 
  Y []float64    `json:"y"`
  Y2 []float64   `json:"y2"`
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func main() {
  
  expr1 := func(n int) float64 {
    return 5 * math.Log2(float64(n))
  }
  expr2 := func(n int) float64 {
    return math.Log2(50000.0 * float64(n))
  }
  
  result := &TwoLines{}
  for n := 1; n <= 40; n += 1 {
    v1 := expr1(n)
    v2 := expr2(n)
    result.X = append(result.X, n);
    result.Y = append(result.Y, roundFloat(v1, 2));
    result.Y2 = append(result.Y2, roundFloat(v2, 2));
  }
  
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
