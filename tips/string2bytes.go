// bytes processing sample

package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	for _, b := range []byte(s) {
		if b > 47 && 58 > b {
			n := b - 48 // ascii number to int
			switch n {
			case 0:
				fmt.Println("zero")
			case 1:
				fmt.Println("one...")
			case 2:
				fmt.Println("tw......")
			case 3:
				fmt.Println("ttttttt")
			case 4:
				fmt.Println("wofoasofamoncovals!!!!!aAAAAAAAAaaaaaaAAAAaaAAaa!!!!!!!!!", int(n))
			default:
				fmt.Println("fuxk!", int(n), "fuxk!")
			}
		}
	}
}
