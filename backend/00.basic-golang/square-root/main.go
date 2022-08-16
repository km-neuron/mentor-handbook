package main

import "fmt"

func main() {
	var a float64
	fmt.Scanf("%f", &a)

	var y float64
	n := 100

	x := 1.0
	for i := 1; i <= n; i++ {
		y = x - (x*x-a)/(2*x)
		x = y
	}

	fmt.Println(y)
}
