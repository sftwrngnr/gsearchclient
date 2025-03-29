package html

import (
	"fmt"
	"html/template"
	"net/http"
)

func SignupPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Retrieve signup form data.
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Perform signup logic here (e.g., store user data in a database).
		// For simplicity, we'll just print the data for demonstration.
		fmt.Printf("New user signup: Username - %s, Password - %s\n", username, password)

		// Redirect to a welcome or login page after signup.
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// If not a POST request, serve the signup page template.
	tmpl, err := template.ParseFiles("html/raw/signup.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

// LoginPage is the handler for the login page.
func LoginPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Perform authentication logic here (e.g., check against a database).
		// For simplicity, we'll just check if the username and password are both "admin".
		if username == "admin" && password == "admin" {
			// Successful login, redirect to a welcome page.
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		// Invalid credentials, show the login page with an error message.
		fmt.Fprintf(w, "Invalid credentials. Please try again.")
		return
	}

	// If not a POST request, serve the login page template.
	tmpl, err := template.ParseFiles("templates/login.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

// WelcomePage is the handler for the welcome page.
func WelcomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome, you have successfully logged in!")
}
