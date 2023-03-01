package main

import "fmt"

func main() {
	var n, number, sum int
	fmt.Print("Enter the number of inputs: ")
	fmt.Scanln(&n)

	fmt.Print("Enter the numbers separated by commas: ")
	for i := 0; i < n; i++ {
		fmt.Scan(&number)
		if i < n-1 {
			fmt.Scan(&struct{}{})
		}
		sum += number
	}

	fmt.Printf("Sum of the numbers: %d\n", sum)
}
