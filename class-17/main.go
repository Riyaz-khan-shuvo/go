package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello I am working!!!")

	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact) // creating the route

	http.ListenAndServe(":8888", nil) //creating the server

}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello I am from go lang and this is my first run code in the web browser")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is about page")
}

func contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello I am from contact page ")
}
