/********************
 * string2slice2.go
 * 2016/11/17(木)
 * walkingmask
 * convert string to slice sample 2
 ********************/

package main

import (
  "fmt"
  "strings"
)

func main() {
  // h: height, w: width
  var h, w int
  fmt.Scan(&h, &w)
  // s: 2-dim string slice
  s := make([][]string, h)
  for i, l := 0, ""; i < h; i++ {
    fmt.Scan(&l)
    s[i] = strings.Split(l, "")
  }
  fmt.Println(s)
}
