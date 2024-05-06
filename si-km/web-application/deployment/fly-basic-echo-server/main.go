package main

import (
	"fmt"
	"net/http"
)

func echoHandler(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Query()["message"][0]
	fmt.Fprintf(w, "Echo: %s", message)
}

func main() {
	http.HandleFunc("/echo", echoHandler)
	fmt.Println("Starting server on port 8080")
	http.ListenAndServe(":8080", nil)
}
