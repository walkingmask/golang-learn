// Command line arguments sample (Simpson's rule)
// integrate 0 to 1 f(x) = 1/(1+x^2) dx

package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	// Get small section width
	args := os.Args
	var m int
	if v, err := strconv.Atoi(args[1]); err != nil {
		fmt.Printf("%q looks like a not number.\n", m)
	} else {
		m = v
	}
	// Function definition
	f := func(x float64) float64 {
		return 1.0 / (1.0 + x*x)
	}
	// Integration interval a to b
	var a, b float64
	a, b = float64(0), float64(1)
	// Calculate
	tval := math.Pi / 4
	res  := Smf(f, a, b, m)
	err  := math.Abs(res - tval)
	rerr := err / math.Abs(tval)
	nsd  := -math.Log10(rerr)
	// Comparison
	fmt.Println("If               =", tval) // Treu value
	fmt.Println("Smf              =", res)  // Simpson's rule
	fmt.Println("|dx|             =", err)  // Abs error
	fmt.Println("|dx|/|x|         =", rerr) // Relative error
	fmt.Println("-log10(|dx|/|x|) =", nsd)  // Number of significant digits
}

func Smf(f func(float64) float64, a, b float64, m int) float64 {
	h, resm1, resm2 := (b-a)/float64(m), 0.0, 0.0
	for i := 1; i < m; i++ {
		resm1 += f(a + float64(i)*h)
	}
	for i := 0; i < m; i++ {
		resm2 += f(a + (float64(i)+0.5)*h)
	}
	return (h / 6) * (f(a) + f(b) + 2.0*resm1 + 4.0*resm2)
}
