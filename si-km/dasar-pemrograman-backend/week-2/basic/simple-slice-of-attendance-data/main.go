package main

import (
	"fmt"
	"strings"
)

func main() {
	attendanceString := "John, Present, Sick, Present, Absent, Mary, Present, Present"
	attendanceArray := strings.Split(attendanceString, ", ")
	var attendanceSlice [][]string
	var currentSlice []string

	for i := 0; i < len(attendanceArray); i++ {
		if i%5 == 0 {
			if currentSlice != nil {
				attendanceSlice = append(attendanceSlice, currentSlice)
			}
			currentSlice = []string{}
		}
		currentSlice = append(currentSlice, attendanceArray[i])
	}
	attendanceSlice = append(attendanceSlice, currentSlice)

	fmt.Println(attendanceSlice)
}
