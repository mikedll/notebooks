package main

import (
  "math"
  "log"
  "encoding/json"
  "io/ioutil"
  "github.com/dougwatson/Go/v3/math/factorial"
  "fmt"
  "strconv"
)

type TwoLines struct {
  X []int64       `json:"x"` 
  Y []float64     `json:"y"`
  Y2 []float64   `json:"y2"`
}

type ThreeLines struct {
  X []int64       `json:"x"` 
  Y []float64     `json:"y"`
  Y2 []float64    `json:"y2"`
  Y3 []float64    `json:"y3"`
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func sciNotation(input float64) string {
  count := 0
  for input >= 10 {
    count += 1
    input = input / 10.0;
  }
  return strconv.FormatFloat(input, 'e', -1, 64) + "e" + strconv.Itoa(count);
}

func exercise3p3d4c() *ThreeLines {
  expr1 := func(n int64) float64 {
    return 5 * math.Log2(float64(n))
  }
  expr2 := func(n int64) float64 {
    return math.Log2(50000.0 * float64(n))
  }
  expr3 := func(n int64) float64 {
    inputN := n + 10;
    if(inputN <= 0) {
      return 0;
    }
    return math.Log2(50000.0 * float64(inputN))
  }
  
  result := &ThreeLines{}
  for n := int64(1); n <= 40; n += 1 {
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
  expr1 := func(n int64) float64 {
    return roundFloat(math.Pow(float64(n), 3), 3);
  }
  
  expr2 := func(n int64) float64 {
    factArg := int(math.Ceil(math.Log2(math.Log2(float64(n)))))
    return roundFloat(float64(factorial.Iterative(factArg)), 3)
  }
  
  result := &TwoLines{}
  for n := int64(1); n <= 30; n++ {
    v1 := expr1(n);
    v2 := expr2(n);
    result.X = append(result.X, n)
    result.Y = append(result.Y, v1)
    result.Y2 = append(result.Y2, v2)
  }
  
  return result;
}

func exercise3p3d5a() *TwoLines {
  expr1 := func(n float64) float64 {
    factArg := int(math.Ceil(math.Log2(n)))
    return roundFloat(float64(factorial.Iterative(factArg)), 3)
  }
  
  expr2 := func(n float64) float64 {
    return roundFloat(math.Pow(n, 2), 3);
  }
    
  
  result := &TwoLines{}
  start := math.Pow(float64(2), 8)
  end := math.Pow(float64(2), 12)
  step := (end - start) / 10

  for n := start; n <= end; n += step {
    v1 := expr1(n);
    v2 := expr2(n);
    result.X = append(result.X, int64(n))
    result.Y = append(result.Y, v1)
    result.Y2 = append(result.Y2, v2)
  }
  
  examined := []float64 { 
    math.Pow(float64(2), 20),
    math.Pow(float64(2), 21),
    math.Pow(float64(2), 23),
  } 
  for _, n := range examined {
    v1 := expr1(n)
    v2 := expr2(n)
    
    
    // s := fmt.Sprintf("|%s|%s|%s|", sciNotation(n), sciNotation(v1), sciNotation(v2))
    s := fmt.Sprintf("|%15.2f|%15.f|%15.f|", n, v1, v2)
    fmt.Println(s)
  }
  
  return result;
}

func main() {
  result := exercise3p3d5a();
  
  output, err := json.Marshal(*result)
  if err != nil {
    log.Fatal(err)
  }

  /*
  var outputPretty []byte
  outputPretty, err = json.MarshalIndent(*result, "", "  ")
  if err != nil {
    log.Fatal(err)
  }
  */
  
  ioutil.WriteFile("data.json", output, 0644)
  log.Println("Wrote data.json")
  // println(string(outputPretty))
}
