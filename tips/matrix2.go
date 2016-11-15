/********************
 * matrix2.go
 * 2016/11/15(火)
 * walkingmask
 * matrix sample 2
 * https://godoc.org/github.com/gonum/matrix/mat64
 * https://www.socketloop.com/tutorials/golang-linear-algebra-and-matrix-calculation-example
 ********************/

package main

import (
  "fmt"
  "github.com/gonum/matrix/mat64"
)

func main() {
  // init
  row1 := []float64{1, 2, 3}
  row2 := []float64{4, 5, 6}
  row3 := []float64{7, 8, 9}
  m := mat64.NewDense(3, 3, nil)
  m.SetRow(0, row1)
  m.SetRow(1, row2)
  m.SetRow(2, row3)

  // print
  fmt.Printf("m :\n%v\n\n", mat64.Formatted(m, mat64.Prefix(""), mat64.Excerpt(0)))

  // print a element
  fmt.Printf("m(1,0): \n%f\n\n", m.At(1, 0))

  // transpose
  mT := m.T()
  fmt.Printf("mT :\n%v\n\n", mat64.Formatted(mT, mat64.Prefix(""), mat64.Excerpt(0)))

  // calc
  mX := mat64.NewDense(3, 3, nil)
  mX.Add(m, mT)
  fmt.Printf("m + mX :\n%v\n\n", mat64.Formatted(mX, mat64.Prefix(""), mat64.Excerpt(0)))
  mX.Sub(m, mT)
  fmt.Printf("m - mX :\n%v\n\n", mat64.Formatted(mX, mat64.Prefix(""), mat64.Excerpt(0)))
  mX.MulElem(m, mT)
  fmt.Printf("m * mX :\n%v\n\n", mat64.Formatted(mX, mat64.Prefix(""), mat64.Excerpt(0)))
  mX.DivElem(m, mT)
  fmt.Printf("m / mX :\n%v\n\n", mat64.Formatted(mX, mat64.Prefix(""), mat64.Excerpt(0)))

  // product
  mX.Mul(m, mT)
  fmt.Printf("m x mX :\n%v\n\n", mat64.Formatted(mX, mat64.Prefix(""), mat64.Excerpt(0)))

  // inverse
  mX.Inverse(m)
  fmt.Printf("m inv :\n%v\n\n", mat64.Formatted(mX, mat64.Prefix(""), mat64.Excerpt(0)))
  mX.Mul(m, mX)
  fmt.Printf("m x inv :\n%v\n\n", mat64.Formatted(mX, mat64.Prefix(""), mat64.Excerpt(0)))

  // determinant
  fmt.Printf("det m :\n%v\n\n", mat64.Det(m))

  // solve
  mX.Solve(m, mX)
  fmt.Printf("Solved m,mX :\n%v\n", mat64.Formatted(mX, mat64.Prefix(""), mat64.Excerpt(0)))
}
