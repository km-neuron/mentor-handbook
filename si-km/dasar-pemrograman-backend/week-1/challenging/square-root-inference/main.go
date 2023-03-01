package main

import (
	"fmt"
)

func main() {
	var a, x float64
	fmt.Print("Enter a number: ")
	fmt.Scan(&a)

	// Initial guess
	x = 1.0

	// Number of iterations
	n := 10

	for i := 0; i < n; i++ {
		x = x - ((x*x - a) / (2 * x))
	}

	fmt.Printf("The estimated square root of %v is %v\n", a, x)
}
