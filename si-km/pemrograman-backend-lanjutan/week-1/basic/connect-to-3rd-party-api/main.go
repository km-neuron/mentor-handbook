package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Name       string `json:"name"`
	Capital    string `json:"capital"`
	Population int    `json:"population"`
}

func main() {
	// Create the client
	client := &http.Client{}

	// Connect to 3rd party API data
	resp, err := client.Get("https://restcountries.com/v2/name/indonesia")
	if err != nil {
		fmt.Println("Error connecting to API:", err)
		return
	}
	defer resp.Body.Close()

	// Decode the JSON response
	var data []Response
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		fmt.Println("Error decoding JSON response:", err)
		return
	}

	// Output the data to terminal
	fmt.Println("Country:", data[0].Name)
	fmt.Println("Capital:", data[0].Capital)
	fmt.Println("Population:", data[0].Population)
}
