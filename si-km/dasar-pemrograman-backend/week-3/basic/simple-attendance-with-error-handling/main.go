package main

import (
	"errors"
	"fmt"
	"strings"
)

type Attendance struct {
	Name    string
	RawData string
}

func (a *Attendance) populate(input string) error {
	data := strings.Split(input, ", ")
	a.Name = data[0]
	for _, d := range data[1:] {
		switch d {
		case "Present":
			a.RawData += "Present, "
		case "Absent":
			a.RawData += "Absent, "
		case "Sick":
			a.RawData += "Sick, "
		case "Not Present":
			return errors.New("Not a valid value")
		default:
			return errors.New("Invalid data")
		}
	}
	a.RawData = strings.TrimSuffix(a.RawData, ", ")
	return nil
}

func main() {
	input := "Budi, Present, Present, Sick, Present, Ani, Present, Present, Absent, Absent, Not Present"

	attendances := []Attendance{}
	inputs := strings.Split(input, ", Not Present")
	for _, i := range inputs {
		a := Attendance{}
		err := a.populate(i)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}
		attendances = append(attendances, a)
	}

	for _, a := range attendances {
		fmt.Printf("Name: %s, Raw Data: %s\n", a.Name, a.RawData)
	}
}
