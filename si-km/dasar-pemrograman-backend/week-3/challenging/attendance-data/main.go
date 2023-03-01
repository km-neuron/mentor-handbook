package main

import (
	"fmt"
	"strings"
)

type Attendance struct {
	Name    string
	RawData string
}

func (a *Attendance) Populate(rawData string) error {
	if strings.Contains(rawData, "Not Present") {
		return fmt.Errorf("Not a valid value")
	}
	a.RawData = rawData
	return nil
}

func main() {
	inputStr := "Budi, Present, Present, Sick, Present, Ani, Present, Present, Absent, Absent"

	// Split the input string into an array of words
	words := strings.Split(inputStr, ", ")

	// Create an array of Attendance structs
	var attendanceArray []Attendance

	// Loop through the array of words and update the attendanceArray accordingly
	var currentAttendance Attendance
	for i, word := range words {
		if i%5 == 0 {
			if i != 0 {
				attendanceArray = append(attendanceArray, currentAttendance)
			}
			currentAttendance = Attendance{Name: word}
		} else {
			err := currentAttendance.Populate(word)
			if err != nil {
				fmt.Println("Error:", err)
			}
		}
	}
	// Append the last Attendance object
	attendanceArray = append(attendanceArray, currentAttendance)

	// Print the resulting array of structs
	fmt.Println(attendanceArray)
}
