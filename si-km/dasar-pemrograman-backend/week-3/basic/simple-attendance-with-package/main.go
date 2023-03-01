package main

import (
	"errors"
	"fmt"
	"time"
)

type Attendance struct {
	Name     string
	RawData  string
	Date     string
	DateTime time.Time
}

func (a *Attendance) populate() error {
	if a.RawData == "" {
		return errors.New("Not a valid value")
	}
	if containsNotPresent(a.RawData) {
		return errors.New("Not a valid value")
	}

	a.Date = time.Now().Format("02 Jan 2006")
	a.DateTime = time.Now()

	return nil
}

func containsNotPresent(rawData string) bool {
	return rawData == "Not Present" || contains(rawData, ", Not Present") || contains(rawData, "Not Present, ")
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && s[len(s)-len(substr):] == substr
}

func main() {
	attendance1 := &Attendance{Name: "Budi", RawData: "Present, Present, Sick, Present"}
	err := attendance1.populate()
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}

	attendance2 := &Attendance{Name: "Ani", RawData: "Present, Present, Absent, Absent, Not Present"}
	err = attendance2.populate()
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}

	fmt.Println(attendance1)
	fmt.Println(attendance2)
}
