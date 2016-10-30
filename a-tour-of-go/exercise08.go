package main

import (
  "io"
  "os"
  "strings"
)

type rot13Reader struct {
  r io.Reader
}

func (rot rot13Reader) Read(b []byte) (int, error) {
  rot.r.Read(b)
  i := 0
  for ; b[i] != 0; i++ {
    switch {
    case b[i] >= 65 && b[i] <= 90:
      b[i] = ((b[i] - 65 + 13) % 26) + 65
    case b[i] >= 97 && b[i] <= 122:
      b[i] = ((b[i] - 97 + 13) % 26) + 97
    }
  }
  return i, io.EOF
}

func main() {
  s := strings.NewReader("Lbh penpxrq gur pbqr!")
  r := rot13Reader{s}
  io.Copy(os.Stdout, &r)
}
