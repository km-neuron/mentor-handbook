package main

import (
	"fmt"
	"net/http"
)

var authenticated bool // flag to check if user is authenticated

func main() {
	http.HandleFunc("/", loginHandler)
	http.HandleFunc("/dashboard", dashboardHandler)

	fmt.Println("Listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// check if username and password are correct
		username := r.FormValue("username")
		password := r.FormValue("password")

		if username == "user" && password == "password" {
			authenticated = true
			http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
			return
		}
	}

	// show login page
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, `
		<h1>Login</h1>
		<form method="post">
			<label>Username:</label>
			<input type="text" name="username" required><br>
			<label>Password:</label>
			<input type="password" name="password" required><br>
			<input type="submit" value="Login">
		</form>
	`)
}

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	if !authenticated {
		// if user is not authenticated, redirect to login page
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// show dashboard
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Hello World!</h1>")
}
