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

	// Create a map to store attendance data for each person
	attendanceData := make(map[string]map[string]int)
	for i := 0; i < len(words); i += 5 {
		name := words[i]
		attendanceData[name] = make(map[string]int)
		attendanceData[name]["Present"] = 0
		attendanceData[name]["Absent"] = 0
		attendanceData[name]["Sick"] = 0
		for j := 1; j <= 4; j++ {
			if words[i+j] == "Present" {
				attendanceData[name]["Present"]++
			} else if words[i+j] == "Absent" {
				attendanceData[name]["Absent"]++
			} else if words[i+j] == "Sick" {
				attendanceData[name]["Sick"]++
			}
		}
	}

	// Print the resulting map
	fmt.Println(attendanceData)
}
