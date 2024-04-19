package main

import (
  "math"
  "log"
)

func main() {
  
  expr1 := func(n int) float64 {
    return math.Log2(50000.0 * float64(n))
  }
  expr2 := func(n int) float64 {
    return 5 * math.Log2(float64(n))
  }
  
  log.Printf("|----------|----------|")
  for n := 10; n <= 20; n += 2 {
    v1 := expr1(n)
    v2 := expr2(n)
    log.Printf("|%10.2f|%10.2f|\n", v1, v2)
  }
  
}
