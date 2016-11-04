// 2 variables for loop sample

package main

import "fmt"

func main() {
  for x, y := 0, 0; x < 10 && y < 10; x, y = x+1, y+1 {
    fmt.Println(x, y)
  }
}
