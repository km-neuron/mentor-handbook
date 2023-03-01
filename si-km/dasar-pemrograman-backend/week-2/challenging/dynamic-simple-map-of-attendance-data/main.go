package main

import (
	"fmt"
	"strings"
)

func main() {
	// Take input from user
	var inputStr string
	fmt.Print("Enter the attendance string: ")
	fmt.Scanln(&inputStr)

	// Split the input string into an array of words
	words := strings.Split(inputStr, ", ")

	// Create a map with keys based on the input string
	attendanceMap := make(map[string]int)
	for _, word := range words {
		if _, ok := attendanceMap[word]; !ok {
			attendanceMap[word] = 0
		}
	}

	// Loop through the array of words and update the map accordingly
	for _, word := range words {
		attendanceMap[word]++
	}

	// Print the resulting map
	fmt.Println(attendanceMap)
}
