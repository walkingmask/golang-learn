/********************
 * file.go
 * 2016/11/18(金)
 * walkingmask
 * https://github.com/astaxie/build-web-application-with-golang/blob/master/ja/07.5.md
 * http://qiita.com/knt45/items/557ee65c46a685ea4f59
 ********************/

package main

import (
  "os"
)

func main() {

  // create file
  fname1 := "go-lang-file-write-test.txt"
  file1, err := os.Create(fname1)
  if err != nil {
    panic(err)
  }
  defer file1.Close()

  // write
  for i := 0; i < 10; i++ {
    file1.WriteString("test1\r\n")
    file1.Write([]byte("test2\r\n"))
  }

  // open file
  fname2 := fname1
  file2, err := os.Open(fname2)
  if err != nil {
    panic(err)
  }
  defer file2.Close()

  // read
  buf := make([]byte, 1024)
  for {
    n, _ := file2.Read(buf)
    if 0 == n {
      break
    }
    os.Stdout.Write(buf[:n])
  }

  // move
  if err := os.Rename(fname2, os.Getenv("HOME")+"/Desktop/"+fname2); err != nil {
    panic(err)
  }
}
