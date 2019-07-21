package main

import (
	"fmt"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello pedram!")
}

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	for k, v := range r.Form {
		fmt.Println("username is", k, "and password is", v)
	}

	fmt.Fprintf(w, "login successful!")
}

func signup(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fmt.Println("username is", r.FormValue("username"))
	fmt.Println("email is", r.FormValue("email"))
	fmt.Println("name is", r.FormValue("name"))
	fmt.Println("password is", r.FormValue("password"))

	fmt.Fprintf(w, "signup successful")
}

func main() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/login", login)
	http.HandleFunc("/signup", signup)

	http.ListenAndServe(":9090", nil)


}