/********************
 * import.go
 * 2016/11/18(金)
 * walkingmask
 * import samples
 ********************/

package main

import (
  . "fmt" // period import
  _ "math" // blank import
)

func main() {
  // there is no need to attach a prefix with period import.
  // but, it is bad code because you don't know which package has the func.
  Println("Hello World!")

  // the below occurs a error because blank import executes only the init func of the package.
  // so, if you want to use the init or to comment-out the code as temporarily, use this.
  //Println(math.Pi)
}
