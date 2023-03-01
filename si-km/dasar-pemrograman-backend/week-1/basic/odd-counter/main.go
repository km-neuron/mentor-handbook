package main

import "fmt"

func main() {
	var n, number, count int

	fmt.Print("Enter the number of inputs: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Print("Enter a number: ")
		fmt.Scan(&number)

		if number%2 == 1 {
			count++
		}
	}

	fmt.Printf("The number of odd numbers: %d\n", count)
}
