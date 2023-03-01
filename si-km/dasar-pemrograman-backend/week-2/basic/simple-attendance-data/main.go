package main

import (
	"fmt"
	"strings"
)

func main() {
	// Input string
	inputStr := "Present, Absent, Sick, Present, Present, Present"

	// Convert input string to array
	inputArr := strings.Split(inputStr, ", ")

	// Print original array
	fmt.Println("Original array:", inputArr)

	// Loop through the array and change "Absent" to "Present"
	for i, val := range inputArr {
		if val == "Absent" {
			inputArr[i] = "Present"
		}
	}

	// Print updated array
	fmt.Println("Updated array:", inputArr)
}
