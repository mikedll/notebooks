package main

import (
  "math"
  "log"
  "encoding/json"
  "fmt"
)

type Line struct {
  X []int      `json:"x"` 
  Y []string   `json:"y"`
  Y2 []string   `json:"y2"`
}

func main() {
  
  expr1 := func(n int) float64 {
    return 5 * math.Log2(float64(n))
  }
  expr2 := func(n int) float64 {
    return math.Log2(50000.0 * float64(n))
  }
  
  result := &Line{}
  for n := 10; n <= 20; n += 2 {
    v1 := expr1(n)
    v2 := expr2(n)
    result.X = append(result.X, n);
    result.Y = append(result.Y, fmt.Sprintf("%.2f", v2));
    result.Y2 = append(result.Y2, fmt.Sprintf("%.2f", v1));
  }

  
  output, err := json.Marshal(*result)
  if err != nil {
    log.Fatal(err)
  }
  
  println(string(output))
}
