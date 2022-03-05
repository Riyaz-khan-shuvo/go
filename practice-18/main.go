package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello I am working")

	http.HandleFunc("/", home)
	http.ListenAndServe(":8888", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I am from practice page ")
}
