package main

import (
  "strings"
  "golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
  dict := map[string]int{}
  for _, word := range strings.Fields(s) {
    _, exist := dict[word]
    if exist == true {
      dict[word]++
    } else {
      dict[word] = 1
    }
  }
  return dict
}

func main() {
  wc.Test(WordCount)
}
