// matrix sample

package main

import "fmt"

func main() {
  matrix := initm(16)
  putm(matrix, 1, 3, 9)
  putm(matrix, 3, 12, 1)
  putm(matrix, 9, 1, 0)
  putm(matrix, 5, 3, 15)
  printm(matrix, 16)
}

// init the matrix and return its pointer
func initm(n int) *[][]int {
  matrix := make([][]int, n)
  for i := 0; i < 16; i++ {
    matrix[i] = make([]int, n)
  }
  return &matrix
}

// put a value to the matrix
func putm(matrix *[][]int, val, x, y int) {
  (*matrix)[y][x] = val
}

// print the matrix
func printm(matrix *[][]int, n int) {
  for i := 0; i < n; i++ {
    for j := 0; j < n; j++ {
      fmt.Print((*matrix)[i][j])
    }
    fmt.Print("\n")
  }
  fmt.Print("\n")
}
