// make array sample

package main

import "fmt"

func main() {
  // an array
  x := make([]int, 10)
  //for i := 0; i < 11; i++ { // error
  for i := 0; i < 10; i++ {
    x[i] = i
  }
  fmt.Println(x)

  // two-dimensional array
  //y := make([]int, 10) // error
  //y := make([][]int, 10, 4) // error
  y := make([][]int, 10)
  for i := 0; i < 10; i++ {
    y[i] = make([]int, 4) // need
    y[i][0] = i+0
    y[i][1] = i+1
    y[i][2] = i+2
    y[i][3] = i+3
  }
  fmt.Println(y)
}
