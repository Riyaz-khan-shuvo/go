package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/features", features)
	http.HandleFunc("/docs", docs)
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8888", nil)
}

func homePage(w http.ResponseWriter, r *http.Request) {

	//pToTem = pointer to template

	pToTem, err := template.ParseFiles("./template/base.gohtml")

	if err != nil {
		fmt.Println(err.Error())
	}
	pToTem.Execute(w, nil)

	fmt.Fprintf(w, "Welcome , to the home page ")
}

func features(w http.ResponseWriter, r *http.Request) {
	pToTem, err := template.ParseFiles("./template/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	pToTem, err = pToTem.ParseFiles("./wpage/features.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	pToTem.Execute(w, nil)
}

// func docs(w http.ResponseWriter, r *http.Request) {
// 	pToTem, err := template.ParseFiles("./template/base.gohtml")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	pToTem, err = pToTem.ParseFiles("./wpage/docs.gohtml")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}

// 	pToTem.Execute(w, nil)
// }

func docs(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("template/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("wpage/docs.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)
	//fmt.Fprintf(w, `welcome`)
}

// func features(w http.ResponseWriter, r *http.Request) {

// 	ptmp, err := template.ParseFiles("template/base.gohtml")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}

// 	ptmp, err = ptmp.ParseFiles("wpage/features.gohtml")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	ptmp.Execute(w, nil)
// 	//fmt.Fprintf(w, `welcome`)
// }
