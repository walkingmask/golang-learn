// string2slice sample

package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "strconv"
)

func main() {
  sc := bufio.NewScanner(os.Stdin)
  sc.Scan()
  s := strings.Split(sc.Text(), " ")
  fmt.Println(s)
  // loop and convert to int
  for i, w := range s {
    v, _ := strconv.Atoi(w)
    fmt.Println(i, v)
  }
}