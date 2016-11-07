// goroutine simple sample

package main

import (
	"fmt"
	"time"
)

func main() {
	var n int
	fmt.Scan(&n)
	sem := make(chan int, n)
	for i := n; i > 0; i-- {
		go func(n int, sem chan int) {
			for i := n; i > 0; i-- {
				fmt.Println(n, i)
				time.Sleep(1 * time.Second)
			}
			sem <- 1
		}(i, sem)
	}
	for i := n; i > 0; i-- {
		<-sem
	}
}
