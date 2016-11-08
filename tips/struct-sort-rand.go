// struct sort sample and rand sample

package main

import (
  "fmt"
  "math/rand"
  "sort"
  "time"
)

// struct definition
type point struct {
  x int
  y int
}
type points []point

func main() {
  // rand's seed
  rand.Seed(time.Now().UnixNano())
  // num of line
  n := 10
  line := make([]points, n)
  // init
  for i := 0; i < n; i++ {
    for j := 0; j < 5; j++ {
      line[i] = append(line[i], point{rand.Intn(100), rand.Intn(100)})
    }
  }
  // print1
  for i := 0; i < n; i++ {
    fmt.Println(i, line[i])
  }
  // sort
  fmt.Println("##### sort #####")
  for i := 0; i < n; i++ {
    sort.Sort(line[i])
  }
  // print2
  for i := 0; i < n; i++ {
    fmt.Println(i, line[i])
  }
}

// sort's interface
func (p points) Len() int {
    return len(p)
}
func (p points) Swap(i, j int) {
    p[i], p[j] = p[j], p[i]
}
func (p points) Less(i, j int) bool {
    return p[i].x < p[j].x
}
