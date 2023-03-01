package main

import (
	"fmt"
	"strings"
)

type AttendanceData struct {
	Name    string
	RawData string
}

type AttendanceDataPopulater interface {
	populate(data string)
}

func (a *AttendanceData) populate(data string) {
	attendance := strings.Split(data, ", ")
	a.Name = attendance[0]
	a.RawData = strings.Join(attendance[1:], ", ")
}

func main() {
	inputStr := "Budi, Present, Present, Sick, Present, Ani, Present, Present, Absent, Absent"

	attendanceDataList := make([]AttendanceData, 0)

	attendanceList := strings.Split(inputStr, ", ")

	for i := 0; i < len(attendanceList); i += 5 {
		data := AttendanceData{}
		data.populate(strings.Join(attendanceList[i:i+5], ", "))
		attendanceDataList = append(attendanceDataList, data)
	}

	fmt.Println(attendanceDataList)
}
