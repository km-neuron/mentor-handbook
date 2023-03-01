package main

import (
	"fmt"
	"strings"
)

func main() {
	soulution1()
	soulution2()
}

func soulution1() {
	attendanceData := "Present, Present, Absent, Present"
	attendanceMap := make(map[string]int)
	attendanceList := strings.Split(attendanceData, ", ")

	for _, status := range attendanceList {
		if status == "Present" {
			attendanceMap["Present"]++
		} else if status == "Absent" {
			attendanceMap["Absent"]++
		} else if status == "Sick" {
			attendanceMap["Sick"]++
		}
	}

	fmt.Println(attendanceMap)
}

func soulution2() {
	// Take input from user
	var inputStr string
	fmt.Print("Enter the attendance string: ")
	fmt.Scanln(&inputStr)

	// Create a map with keys 'Present', 'Absent', 'Sick'
	attendanceMap := make(map[string]int)
	attendanceMap["Present"] = 0
	attendanceMap["Absent"] = 0
	attendanceMap["Sick"] = 0

	// Split the input string into an array of words
	words := strings.Split(inputStr, ", ")

	// Loop through the array of words and update the map accordingly
	for _, word := range words {
		if word == "Present" {
			attendanceMap["Present"]++
		} else if word == "Absent" {
			attendanceMap["Absent"]++
		} else if word == "Sick" {
			attendanceMap["Sick"]++
		}
	}

	// Print the resulting map
	fmt.Println(attendanceMap)
}
