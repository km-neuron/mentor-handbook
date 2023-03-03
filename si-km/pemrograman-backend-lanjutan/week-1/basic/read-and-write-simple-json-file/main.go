package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Birthdate string `json:"birthdate"`
}

func main() {
	// Read the file
	file, err := ioutil.ReadFile("json/data.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	// Convert JSON to struct
	var person Person
	err = json.Unmarshal(file, &person)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		os.Exit(1)
	}

	// Output the JSON struct to terminal
	fmt.Println("Current Data:")
	fmt.Println(person)

	// Append a new JSON data to struct
	newPerson := Person{FirstName: "Dito", LastName: "Prasetyo", Birthdate: "1995-01-01"}
	personList := []Person{person, newPerson}

	// Write the new JSON to file
	newFile, err := json.MarshalIndent(personList, "", "    ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		os.Exit(1)
	}

	err = ioutil.WriteFile("json/data.json", newFile, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		os.Exit(1)
	}

	fmt.Println("New Data:")
	fmt.Println(personList)
}
