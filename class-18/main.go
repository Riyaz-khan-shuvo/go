package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", homePage)
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8888", nil)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome , to the home page ")
}
