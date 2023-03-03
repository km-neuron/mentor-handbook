package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	// Parse the HTML template
	tmpl := template.Must(template.ParseFiles("template/index.html"))

	// Serve the HTML page
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func calculate(w http.ResponseWriter, r *http.Request) {
	// Get the form values
	firstStr := r.FormValue("first")
	secondStr := r.FormValue("second")
	operation := r.FormValue("operation")

	// Convert the values to float64
	first, err := strconv.ParseFloat(firstStr, 64)
	if err != nil {
		fmt.Fprintln(w, "Error converting first value to float:", err)
		return
	}

	second, err := strconv.ParseFloat(secondStr, 64)
	if err != nil {
		fmt.Fprintln(w, "Error converting second value to float:", err)
		return
	}

	// Perform the calculation
	var result float64
	switch operation {
	case "+":
		result = first + second
	case "-":
		result = first - second
	case "*":
		result = first * second
	case "/":
		result = first / second
	default:
		fmt.Fprintln(w, "Invalid operation")
		return
	}

	// Print the result to the web page
	fmt.Fprintf(w, "%.2f %s %.2f = %.2f", first, operation, second, result)
}

func main() {
	// Register the routes
	http.HandleFunc("/", homePage)
	http.HandleFunc("/calculate", calculate)

	// Start the web server
	fmt.Println("Starting server on port 8080...")
	http.ListenAndServe(":8080", nil)
}
