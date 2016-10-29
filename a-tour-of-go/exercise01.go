package main

import (
  "fmt"
  "math"
)

func Sqrtf1(x float64) float64 {
  z := x
  for i := 0; i < 10; i++ {
    z = z - ((z*z - x) / (2 * z))
  }
  return z
}

func Sqrtf2(x float64) (float64, int) {
  zb := x
  cnt := 0
  for {
    cnt++
    za := zb - ((zb*zb - x) / (2 * zb))
    if math.Abs(za-zb) < 0.000000000001 {
      break
    } else {
      zb = za
    }
  }
  return zb, cnt
}

func main() {
  for i := 1; i <= 10; i++ {
    fmt.Println("case:", i)
    fmt.Println(" math.Sqrt:  ", math.Sqrt(float64(i)))
    fmt.Println(" Sqrtf1:   ", Sqrtf1(float64(i)), "( 10 )")
    abs, cnt := Sqrtf2(float64(i))
    fmt.Println(" Sqrtf2:   ", abs, "(", cnt, ")")
  }
}
