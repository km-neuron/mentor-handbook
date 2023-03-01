package main

import "fmt"

func main() {
	var name string
	var gender string
	var age int

	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)

	fmt.Print("Enter your gender (male/female): ")
	fmt.Scanln(&gender)

	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)

	title := ""
	if gender == "male" {
		if age >= 17 {
			title = "Mr"
		} else {
			title = "Boy"
		}
	} else if gender == "female" {
		if age >= 17 {
			title = "Ms"
		} else {
			title = "Girl"
		}
	}

	fmt.Printf("Hi %s %s\n", title, name)
}
