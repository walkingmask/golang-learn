package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
  var picture [][]uint8
  for y := 0; y < dy; y++ {
    var row []uint8
    for x := 0; x < dx; x++ {
      row = append(row, uint8(x*y))
    }
    picture = append(picture, row)
  }
  return picture
}

func main() {
  pic.Show(Pic)
}
