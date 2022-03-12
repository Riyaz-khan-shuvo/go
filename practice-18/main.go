package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	fmt.Println("Hello I am working!!!")

	http.HandleFunc("/", homePage)
	http.HandleFunc("/features", featurePage)
	http.HandleFunc("/docs", docsPage)
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("./loruki"))))

	http.ListenAndServe(":8888", nil)

}

func homePage(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hello I am from home page ")
	pToTem, err := template.ParseFiles("./template/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	pToTem.Execute(w, nil)
}

func featurePage(w http.ResponseWriter, r *http.Request) {
	pToTem, err := template.ParseFiles("./template/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	pToTem, err = pToTem.ParseFiles("./wpages/feature.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	pToTem.Execute(w, nil)
}

func docsPage(w http.ResponseWriter, r *http.Request) {
	pToTemp, err := template.ParseFiles("./template/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	pToTemp, err = pToTemp.ParseFiles("./wpages/docs.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	pToTemp.Execute(w, nil)
}

// func about(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "I am from about page ")
// }
