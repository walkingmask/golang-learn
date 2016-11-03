// round float to int with the first decimal place sample

package main

import "fmt"

func round(x float64) int {
  if int(x*10)%10 < 5 {
    return int(x)
  }
  return int(x)+1
}

func main() {
  fmt.Println(round(23.4)) // 23
  fmt.Println(round(45.6)) // 46
}