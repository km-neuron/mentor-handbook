package main

import "fmt"

func main() {
	var name, gender string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	fmt.Print("Enter your gender (male/female): ")
	fmt.Scanln(&gender)

	title := ""
	if gender == "male" {
		title = "Mr"
	} else if gender == "female" {
		title = "Mrs"
	}

	fmt.Printf("Hi %s %s\n", title, name)
}
