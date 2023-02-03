package main

import "fmt"

func main() {
	var name string

	fmt.Print("Tulis namamu? ")
	fmt.Scan(&name)
	fmt.Println("Hi", name)
}
