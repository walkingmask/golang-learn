/********************
 * c.go
 * 2016/11/26(土)
 * walkingmask
 ********************/

package main

/*
#include <stdio.h>
void pri() {
    printf("pripri\n");
}
*/
import "C"

func main() {
  C.pri()
}
