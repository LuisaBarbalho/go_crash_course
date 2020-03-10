package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello World</h1>")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>About</h1>")
}

func main() {
	// like a route

	// localhost:3000
	http.HandleFunc("/", index)
	// localhost:3000/about
	http.HandleFunc("/about", about)
	fmt.Println("Server Starting... ")
	http.ListenAndServe(":3000", nil)
}
