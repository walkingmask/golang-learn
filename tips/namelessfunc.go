// nameless function samples

package main

import "fmt"

func main() {
  // nameless
  f := func(x int) int {
    return x*x
  }
  fmt.Println(f(10))
  // take a function as an argument
  g := func(f func(int) int, x int) int {
    return f(x*x)
  }
  fmt.Println(g(f, 10))
  // take a nameless function
  fmt.Println(g(func(x int) int {
    return x<<10
  }, 10))
}
