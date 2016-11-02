// Standard input sample

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(n)

	var m int
	fmt.Scanf("%d %d", &n, &m)
	fmt.Println(n, m)

	for i := 0; i < 10; i++ {
		fmt.Scanf("%d %d", &n, &m)
		fmt.Println(n, m)
	}
}
