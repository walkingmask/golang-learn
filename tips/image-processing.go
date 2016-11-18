/********************
 * image-processing.go
 * 2016/11/18(金)
 * walkingmask
 * image-processing sample
 * http://d.hatena.ne.jp/taknb2nch/20131231/1388500659
 * http://stackoverflow.com/questions/36573413/change-color-of-a-single-pixel-go-lang-image
 * http://stackoverflow.com/questions/37635426/go-image-manipulation
 ********************/

package main

import (
  //"fmt"
  "image"
  "image/color"
  "image/jpeg"
  "math"
  "net/http"
  "os"
)

func main() {

  // get image from web
  var url string = "http://0117893.lolipop.jp/wallpaper/2013/AKB48/01171440_AKB48_58.jpg"
  response, err := http.Get(url)
  if err != nil {
      panic(err)
  }
  defer response.Body.Close()

  // decode image
  img, _, err := image.Decode(response.Body)
  if err != nil {
      panic(err)
  }

  // processing here
  bounds := img.Bounds()                            // get bounds of image
  nimg := image.NewRGBA(bounds)                     // new image placeholder as a RGBA
  for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
    for x := bounds.Min.X; x < bounds.Max.X; x++ {
      r, g, b, a := img.At(x, y).RGBA()             // get image's r,g,b,a
      r, g, b, a = r>>8, g>>8, b>>8, a>>8           // conv 0..ffff to 0..ff using bit shift
      nimg.Set(x, y, color.RGBA{f(r), f(g), f(b), uint8(a)}) // set gamma'ed r,g,b,a to new image
    }
  }

  // save to file
  file, err := os.Create("./output2.jpg")
  if err != nil {
    panic(err)
  }
  opt := &jpeg.Options{Quality: 100}
  if err = jpeg.Encode(file, nimg, opt); err != nil {
    panic(err)
  }
  defer file.Close()
}

// gamma correction
func f(x uint32) uint8 {
  gamma := 2.0
  return uint8( 255 * math.Pow(float64(x)/255, 1.0/gamma) );
}