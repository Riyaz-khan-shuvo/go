package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello I am working")
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.ListenAndServe(":8888", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello this is our first practice")
	fmt.Fprintf(w, "Hello this is our first practice")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I am from about page")
}
