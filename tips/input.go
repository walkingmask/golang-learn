// Standard input samples

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// a number
	var n int
	fmt.Scan(&n)
	fmt.Println(n)

	// two numbers separated by a space
	var m int
	fmt.Scanf("%d %d", &n, &m)
	fmt.Println(n, m)

	// loop
	for i := 0; i < 3; i++ {
		fmt.Scanf("%d %d", &n, &m)
		fmt.Println(n, m)
	}

	// string
	var s string
	fmt.Scanf("%s", &s)
	fmt.Println(s)

	// string include some spaces
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	fmt.Println(sc.Text())
}
